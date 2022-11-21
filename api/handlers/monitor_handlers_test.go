package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zob456/snapshot/api/models"
	"github.com/zob456/snapshot/api/utils"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbName   = "postgres"
	sslMode  = "disable"
)

var db = utils.ConnectDB(host, port, user, password, dbName, sslMode)

var app = fiber.New()

// device Creates mock networkDevice for assert.Equal test for Test_GetNetworkDevice_HappyPath
var device = models.NetworkDevice{
	MachineID: uuid.MustParse("fa52fe06-35bc-499e-8103-8c980b3437f2"),
	Status: models.Status{
		CpuTemp:  120,
		FanSpeed: 900,
		HddSpace: 256,
	},
	LastLoggedIn: "admin",
	SysTime:      time.Date(2001, time.February, 16, 20, 38, 40, 0, time.UTC),
}

//// postDevice for PostNetworkDevice happy path test
//var postDevice = models.NetworkDevice{
//	MachineID: uuid.MustParse("da52fe06-35bc-499e-8103-8c980b3437f2"),
//	Status: models.Status{
//		CpuTemp:  120,
//		FanSpeed: 900,
//		HddSpace: 256,
//	},
//	LastLoggedIn: "admin",
//	SysTime:      time.Date(2001, time.February, 16, 20, 38, 40, 0, time.UTC),
//}


// GetAllNetworkDevice happy path
func Test_GetAllNetworkDevice_HappyPath(t *testing.T) {
	var expectedRes []models.NetworkDevice
	app.Get("/device", GetAllNetworkDevice(db))
	req := httptest.NewRequest(http.MethodGet, "/device", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(body, &expectedRes)
	if err != nil {
		t.Error(err)
	}
	assert.IsType(t, []models.NetworkDevice{}, expectedRes, fmt.Sprintf("type %v", reflect.TypeOf(expectedRes)))
}

// GetAllNetworkDevice sad paths
func Test_GetAllNetworkDevice_SadPath_BadEP(t *testing.T) {
	app.Get("/device", GetAllNetworkDevice(db))
	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 404, resp.StatusCode)
}

// Test_GetNetworkDevice_HappyPath happy path
func Test_GetNetworkDevice_HappyPath(t *testing.T) {
	var actualRes models.NetworkDevice
	app.Get("/device/:id", GetNetworkDevice(db))
	req := httptest.NewRequest(http.MethodGet, "/device/fa52fe06-35bc-499e-8103-8c980b3437f2", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(body, &actualRes)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, device, actualRes)
}

// Test_GetNetworkDevice_HappyPath sad paths
func Test_GetNetworkDevice_SadPath_BadID_Format(t *testing.T) {
	app.Get("/device/:id", GetNetworkDevice(db))
	req := httptest.NewRequest(http.MethodGet, "/device/fa52fe06-35bc", nil)
	resp, err := app.Test(req)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	assert.Equal(t, 400, resp.StatusCode)
}

func Test_GetNetworkDevice_SadPath_Nonexistent_ID(t *testing.T) {
	app.Get("/device/:id", GetNetworkDevice(db))
	req := httptest.NewRequest(http.MethodGet, "/device/fa52fe06-35bc-499e-8103-8c980b3437f4", nil)
	resp, err := app.Test(req)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	assert.Equal(t, 404, resp.StatusCode)
}

// PostNetworkDevice happy path
func Test_PostNetworkDevice_HappyPath(t *testing.T) {
	app.Post("/device/new", PostNetworkDevice(db))
	mockDevice := models.NetworkDevice{
		MachineID: uuid.New(),
		Status: models.Status{
			CpuTemp:  120,
			FanSpeed: 900,
			HddSpace: 256,
		},
		LastLoggedIn: "admin",
		SysTime:      time.Date(2001, time.February, 16, 20, 38, 40, 0, time.UTC),
	}
	reqBody, err := json.Marshal(mockDevice)
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/device/new", bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	assert.Equal(t, 201, resp.StatusCode)
}

// PostNetworkDevice sad path MachineID already exists
func Test_PostNetworkDevice_SadPath_IDExists(t *testing.T) {
	app.Post("/device/new", PostNetworkDevice(db))
	mockDevice := models.NetworkDevice{
		MachineID: uuid.MustParse("fa52fe06-35bc-499e-8103-8c980b3437f2"),
		Status: models.Status{
			CpuTemp:  120,
			FanSpeed: 900,
			HddSpace: 256,
		},
		LastLoggedIn: "admin",
		SysTime:      time.Date(2001, time.February, 16, 20, 38, 40, 0, time.UTC),
	}
	reqBody, err := json.Marshal(mockDevice)
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/device/new", bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	assert.Equal(t, 422, resp.StatusCode)
}

// PostNetworkDevice sad path bad payload
func Test_PostNetworkDevice_SadPath_BadPayload(t *testing.T) {
	app.Post("/device/new", PostNetworkDevice(db))
	mockDevice := models.NetworkDevice{
		MachineID: uuid.New(),
		Status: models.Status{
			CpuTemp:  120,
			FanSpeed: 900,
			HddSpace: 256,
		},
		LastLoggedIn: "admin",
	}

	reqBody, err := json.Marshal(mockDevice)
	if err != nil {
		log.Println(err)
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/device/new", bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 400, resp.StatusCode)
}
