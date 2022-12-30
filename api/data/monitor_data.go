package data

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/zob456/snapshot/api/models"
	"log"
)

func SelectNetworkDeviceData(db *sql.DB, machineID uuid.UUID) (*models.NetworkDevice, error) {
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
		FROM "Network"."vw_device"
		WHERE "ID" = $1;`
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

func SelectAllNetworkDeviceData(db *sql.DB) ([]*models.NetworkDevice, error) {
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

	var networkDevices []*models.NetworkDevice
	rows, err := PostgresScanRows(db, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		networkDevice := &models.NetworkDevice{}
		err = rows.Scan(
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
		networkDevices = append(networkDevices, networkDevice)
	}
	return networkDevices, nil
}

func CreateNetworkDevice(db *sql.DB, device models.NetworkDevice) error {
	// comment below is used by Intellij IDEA to provide sql syntax support & sql verification from the data source
	/*language=PostgreSQL*/
	const deviceQuery = `
		INSERT INTO "Network".device (
			"ID",
			last_logged_in,
		  	sys_time
	  	) 
		VALUES ($1, $2, $3);`

	err := PostgresUpdateColumnDataOneRow(db, deviceQuery, device.MachineID, device.LastLoggedIn, device.SysTime)
	if err != nil {
		return err
	}

	// comment below is used by Intellij IDEA to provide sql syntax support & sql verification from the data source
	/*language=PostgreSQL*/
	const statusQuery = `
		INSERT INTO "Network".status (
			"machine_ID",
			cpu_temp,
		  	hdd_space,
		  	fan_speed
	  	) 
		VALUES ($1, $2, $3, $4);`

	err = PostgresUpdateColumnDataOneRow(db, statusQuery, device.MachineID, device.CpuTemp, device.HddSpace, device.FanSpeed)
	if err != nil {
		secondaryErr := statusFailureRollback(db)
		if secondaryErr != nil {
			log.Printf("ERROR: FAILED to rollback network device data after status table failed to update for device ID: %s", device.MachineID)
		}
		return err
	}

	return nil
}

func statusFailureRollback(db *sql.DB) error {
	// comment below is used by Intellij IDEA to provide sql syntax support & sql verification from the data source
	/*language=PostgreSQL*/
	const deviceQuery = `
		DELETE FROM "Network".device
		WHERE "ID" = $1;`

	err := PostgresUpdateColumnDataOneRow(db, deviceQuery)
	if err != nil {
		return err
	}
	return nil
}