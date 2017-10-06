package main

import(
	"fmt"
	"os"
	"math"
	"bufio"
	"strconv"
)

type counties struct {
	lat float64
	long float64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func haversine(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
func distance(county1 counties, county2 counties) float64 {
	var lat_1, long_1, lat_2, long_2, r float64
	lat_1 = county1.lat * math.Pi / 180
	long_1 = county1.long * math.Pi / 180
	lat_2 = county2.lat * math.Pi / 180
	long_2 = county2.long * math.Pi / 180

	r = 3959 //radius of earth in miles

	h := haversine(lat_2-lat_1) + math.Cos(lat_1)*math.Cos(lat_2)*haversine(long_2-long_1)

	return 2 * r * math.Asin(math.Sqrt(h))

}
func input(county_info map[string]counties) float64{
	var county1 counties
	var county2 counties
	var distance_sum float64

	distance_sum = 0.0

	fmt.Println("First County Name: ")

	input_scan := bufio.NewScanner(os.Stdin)
	var input_line string
	input_line = "stuff"

	county1.lat = 200
	county2.lat = 200

	for input_scan.Scan() {
		input_line = input_scan.Text()

		fmt.Println(input_line)

		if input_line == "" {
			return distance_sum
		}

		if value, ok := county_info[input_line]; ok { 
			if county1.lat == 200 {
			county1 = value 
			} else { if county2.lat == 200 {
			county2 = value
			}
			}
		fmt.Println(county1)
		fmt.Println(county2)
	

		}
	

		if county1.lat != 200 && county2.lat != 200{
			distance_sum = distance_sum + distance(county1, county2)

			county1 = county2
			county2.lat = 200

		}
		
	}	
	return 1
}

func main() {
	var county_info map[string]counties
	var coords counties
	var name string
	var lat_d string
	var long_d string
	county_info = make(map[string]counties)

	f, err := os.Open("micount.txt")
	check(err)

	defer f.Close()

	f_scan := bufio.NewScanner(f)
	f_scan.Split(bufio.ScanWords)

	a := 1

	for f_scan.Scan() {
		if a % 3 == 1 {
			name = f_scan.Text()
		}

		if a % 3 == 2 {
			lat_d = f_scan.Text()
		}

		if a % 3 == 0 {
			long_d = f_scan.Text()
			coords.lat, _ = (strconv.ParseFloat(lat_d, 64))
			coords.long, _ = (strconv.ParseFloat(long_d, 64))
			county_info[name] = coords

		}
		a = a+1
	}

	for k, v := range county_info{
		fmt.Println(k,v)
	}

	d := (input(county_info))
	fmt.Println(d)
}
