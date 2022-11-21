# Snapshot - GoLang REST API for network machine checking

### To verify that you have the tools needed for this project / install them:
Please run `bin/init.sh` to make sure you have all the necessary software installed. (script will install any missing critical software)

This project is way over-engineered beyond what something this simple should be.
I decided to go overboard by creating wrapper functions & distinct error handlers to show more detail around my thought process around engineering. 
I am hoping this gives you a much clearer picture of my style & thought process around coding in general than if I had kept it as simple as it should've been

#### Start the application:
1. Run the command `tilt up` (add `--stream` flag if you prefer to see the log output in the console)


#### Stop the application 
1. `CTRL + c`
2. Run the command `tilt down`. 
   1. **NOTE** make sure to remove the persistent volumes with the command `kubectl delete pvc --all`

#### Run tests:
1. Run the command `tilt up snapshotdb` to make sure the Postgres DB is running
2. Open a new terminal window
3. Run the command `go test ./api/...`
4. **NOTE** Make sure to shut down the DB instance & clean the pvc 
   1. Run the command `tilt down`.
   2. Run the command `kubectl delete pvc --all`


_You can find a JSON file you can use in [Insomnia](https://insomnia.rest/download) (It should also work for Postman) [here](api_tools_json/Insomnia_requests.json). It has a collection that can be used to run requests against the running service_