-- Create schema(s)

CREATE SCHEMA IF NOT EXISTS "Network";

SET SCHEMA 'Network';

CREATE TABLE IF NOT EXISTS "device"
(
    "ID"             uuid        NOT NULL PRIMARY KEY,
    last_logged_in text        NOT NULL,
    sys_time       timestamptz NOT NULL DEFAULT now()
);


-- Status data for network device
CREATE TABLE IF NOT EXISTS "status"
(
    "machine_ID" uuid    NOT NULL PRIMARY KEY REFERENCES "Network".device ("ID"),
    cpu_temp   integer NOT NULL,
    hdd_space  integer NOT NULL,
    fan_speed  integer NOT NULL
);