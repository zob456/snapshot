CREATE OR REPLACE VIEW "Network"."vw_device" as
(
SELECT "ID",
       cpu_temp,
       hdd_space,
       fan_speed,
       last_logged_in,
       sys_time
FROM "Network".device d
         LEFT JOIN "Network".status s on s."machine_ID" = d."ID"
    );