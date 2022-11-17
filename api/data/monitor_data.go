package data

import (
	"database/sql"

	"github.com/zob456/snapshot/api/models"
)

func SelectNetworkDeviceData(db *sql.DB, machineID string) (*models.NetworkDevice, error) {

	// comment below is used by Intellij IDEA to provide sql syntax support & sql verification from the data source
	/*language=PostgreSQL*/
	const query = `
		SELECT
			"ID",
			"cpu_temp",
			"fan_speed",
			"hdd_space",
			"last_logged_in",
			"sys_time"
		FROM "Network"."vw_device";`

	networkDevice := &models.NetworkDevice{}
	row, err := PostgresScanOneRow(db, query, machineID)
	if err != nil {
		return nil, err
	}
	err = row.Scan(
		&networkDevice.MachineID,
		&networkDevice.Status.CpuTemp,
		&networkDevice.Status.FanSpeed,
		&networkDevice.Status.HddSpace,
		&networkDevice.LastLoggedIn,
		&networkDevice.SysTime,
	)
	if err != nil {
		return nil, err
	}
	return networkDevice, nil
}

func SelectAllNetworkDeviceData(db *sql.DB, machineID string) (*models.NetworkDevice, error) {

	// comment below is used by Intellij IDEA to provide sql syntax support & sql verification from the data source
	/*language=PostgreSQL*/
	const query = `
		SELECT
			"ID",
			"cpu_temp",
			"fan_speed",
			"hdd_space",
			"last_logged_in",
			"sys_time"
		FROM "Network"."vw_device";`

	networkDevice := &models.NetworkDevice{}
	row, err := PostgresScanOneRow(db, query, machineID)
	if err != nil {
		return nil, err
	}
	err = row.Scan(
		&networkDevice.MachineID,
		&networkDevice.Status.CpuTemp,
		&networkDevice.Status.FanSpeed,
		&networkDevice.Status.HddSpace,
		&networkDevice.LastLoggedIn,
		&networkDevice.SysTime,
	)
	if err != nil {
		return nil, err
	}
	return networkDevice, nil
}