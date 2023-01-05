//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 3
	Retired     = 2
)

func printServer(listOfServer map[string]int) {
	fmt.Println("the number of servers is", len(listOfServer))
	totalOnline := 0
	totalOffLine := 0
	totalMaintenance := 0
	totalRetired := 0

	for _, server := range listOfServer {
		switch server {
		case Online:
			totalOnline += 1
		case Offline:
			totalOffLine += 1
		case Maintenance:
			totalMaintenance += 1
		case Retired:
			totalRetired += 1
		}
	}

	fmt.Println("total number of servers online", totalOnline)
	fmt.Println("total number of servers offline", totalOffLine)
	fmt.Println("total number of servers under maintenance", totalMaintenance)
	fmt.Println("total number of servers retired", totalRetired)

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverMap := map[string]int{}

	for _, server := range servers {
		serverMap[server] = Online
	}

	printServer(serverMap);
	serverMap["darkstar"] = Retired;
	serverMap["aiur"] = Offline;
	fmt.Println("servers info after changes");
	fmt.Println("**************************");
	printServer(serverMap);

	for key := range serverMap {  
		keyValue := key
		serverMap[keyValue] = Maintenance
	}

	fmt.Println("after new changes");
	fmt.Println("**************************")

    printServer(serverMap)

}
