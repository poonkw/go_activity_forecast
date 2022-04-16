package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Activity struct which contains
// a name and a outdoor
type Activity struct {
	Name    string `json:"name"`
	Outdoor string `json:"outdoor"`
}

var outdoorConds []int = []int{800, 801, 802, 803, 804, 951, 952, 953, 954, 955, 956}

func main() {
	reader := bufio.NewReader(os.Stdin)
	option := startPlanner(reader)
	runOptions(option, reader)
}

func startPlanner(r *bufio.Reader) int {

	fmt.Println("Welcome to your day planner. Please select the following:")
	fmt.Println("1) Show me what I can do today!")
	fmt.Println("2) Add an interesting activity")
	fmt.Println("3) Remove an activity for the foreseeable future")

	input, _ := r.ReadString('\n')
	i, err := strconv.Atoi(input[0:1])
	if err != nil || i > 3 || i < 1 {
		fmt.Println("Option unavailable.")
		i = startPlanner(r)
	}
	return i
}

func runOptions(o int, r *bufio.Reader) {
	activities := openJson()
	switch o {
	case 1:
		wCode := getWeather()
		fmt.Println("\nRecommended activites:")
		for _, v := range outdoorConds {
			if v == wCode {
				for _, v := range activities {
					fmt.Println(v.Name)
				}
				fmt.Println("")
				return
			}
		}
		for _, v := range activities {
			if v.Outdoor == "n" {
				fmt.Println(v.Name)
			}
		}
		fmt.Println("")
	case 2:
		fmt.Println("\nPlease input the activity:")
		activity, _ := r.ReadString('\n')
		fmt.Println("Is it an outdoor activity? (y/n)")
		outdoor, _ := r.ReadString('\n')
		if strings.Contains(outdoor, "y") || strings.Contains(outdoor, "n") {
			addActivity(activity, outdoor, activities)
			fmt.Println("")
		} else {
			fmt.Println("Option invalid.")
			runOptions(o, r)
		}
	case 3:
		fmt.Println("")
		for i, v := range activities {
			fmt.Printf("%d) %s\n", i+1, v.Name)
		}
		fmt.Println("Please input the activity number to delete:")
		input, _ := r.ReadString('\n')
		i, err := strconv.Atoi(strings.TrimSuffix(input, "\r\n"))
		if err != nil || i > len(activities) || i < 1 {
			fmt.Println("Option unavailable.")
			runOptions(o, r)
		} else {
			delActivity(i-1, activities)
			fmt.Println("")
		}
	}
}

func openJson() []Activity {
	// Open our jsonFile
	jsonFile, err := os.Open("./activities.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Activities array
	var activities []Activity

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'activities' which we defined above
	json.Unmarshal(byteValue, &activities)

	return activities
}

func addActivity(a string, o string, acts []Activity) {
	acts = append(acts, Activity{Name: strings.TrimSuffix(a, "\r\n"), Outdoor: strings.TrimSuffix(o, "\r\n")})

	file, _ := json.MarshalIndent(acts, "", " ")

	_ = ioutil.WriteFile("./activities.json", file, 0644)
}

func delActivity(i int, a []Activity) {
	a = append(a[:i], a[i+1:]...)

	file, _ := json.MarshalIndent(a, "", " ")

	_ = ioutil.WriteFile("./activities.json", file, 0644)
}
