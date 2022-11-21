INSERT INTO "Network".device ("ID", last_logged_in, sys_time)
VALUES ('FA52FE06-35BC-499E-8103-8C980B3437F2', 'admin', timestamptz('2001-02-16 20:38:40+00')),
       ('2FE6C725-1A42-4227-BD1B-26BFE98D35C6', 'barry', timestamptz('2001-02-16 20:38:40+00')),
       ('950369A4-1854-42C0-89C6-61ADCA8B276A', 'zob', timestamptz('2001-02-16 20:38:40+00'));

INSERT INTO "Network".status ("machine_ID", cpu_temp, hdd_space, fan_speed)
VALUES ('FA52FE06-35BC-499E-8103-8C980B3437F2', 120, 256, 900),
       ('2FE6C725-1A42-4227-BD1B-26BFE98D35C6', 98, 512, 1200),
       ('950369A4-1854-42C0-89C6-61ADCA8B276A', 200, 172, 875);