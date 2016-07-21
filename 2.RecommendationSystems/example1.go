package main

import (
	"fmt"
	"math"
)

type Rating struct {
	movies map[string]float64
	distance float64
}

var users map[string]Rating

func initDataForUsers () {
	users = map[string]Rating {
		"Angelica" : Rating{
			movies: map[string]float64 {
				"Blues Traveler" : 3.5,
				"Broken Bells" : 2.0,
				"Norah Jones" : 4.5,
				"Phoenix" : 5.0,
				"Slightly Stoopid" : 1.5,
				"The Strokes" : 2.5,
				"Vampire Weekend" : 2.0,
			},
		},
		"Bill" : Rating{
			movies: map[string]float64{
				"Blues Traveler": 2.0,
				"Broken Bells": 3.5,
				"Deadmau5": 4.0,
				"Phoenix": 2.0,
				"Slightly Stoopid": 3.5,
				"Vampire Weekend": 3.0,
			},
		},
		"Chan" : Rating{
			movies: map[string]float64{
				"Blues Traveler": 5.0,
				"Broken Bells": 1.0,
				"Deadmau5": 1.0,
				"Norah Jones": 3.0,
				"Phoenix": 5,
				"Slightly Stoopid": 1.0,
			},
		},
		"Dan" : Rating{
			movies: map[string]float64{
				"Blues Traveler": 3.0,
				"Broken Bells": 4.0,
				"Deadmau5": 4.5,
				"Phoenix": 3.0,
				"Slightly Stoopid": 4.5,
				"The Strokes": 4.0,
				"Vampire Weekend": 2.0,
			},
		},
		"Hailey" : Rating{
			movies: map[string]float64{
				"Broken Bells": 4.0,
				"Deadmau5": 1.0,
				"Norah Jones": 4.0,
				"The Strokes": 4.0,
				"Vampire Weekend": 1.0,
			},
		},
		"Jordyn" : Rating{
			movies: map[string]float64{
				"Broken Bells": 4.5,
				"Deadmau5": 4.0,
				"Norah Jones": 5.0,
				"Phoenix": 5.0,
				"Slightly Stoopid": 4.5,
				"The Strokes": 4.0,
				"Vampire Weekend": 4.0,
			},
		},
		"Sam" : Rating{
			movies: map[string]float64{
				"Blues Traveler": 5.0,
				"Broken Bells": 2.0,
				"Norah Jones": 3.0,
				"Phoenix": 5.0,
				"Slightly Stoopid": 4.0,
				"The Strokes": 5.0,
			},
		},
		"Veronica" : Rating{
			movies: map[string]float64{
				"Blues Traveler": 3.0,
				"Norah Jones": 5.0,
				"Phoenix": 4.0,
				"Slightly Stoopid": 2.5,
				"The Strokes": 3.0,
			},
		},
	}
}

func pearson(r1 Rating, r2 Rating) float64 {
	var sum_xy float64
	var sum_x float64
	var sum_y float64
	var sum_x2 float64
	var sum_y2 float64
	var n float64 = 0
	for k, v := range r1.movies {
		for k2, v2 := range r2.movies {
			if k == k2 && v != 0 && v2 != 0 {
				n++;
				sum_xy += v * v2
				sum_x += v
				sum_y += v2
				sum_x2 += v * v
				sum_y2 += v2 * v2
			}
		}
	}

	if n == 0 {
		return 0
	}

	denominator := math.Sqrt(sum_x2 - ((sum_x * sum_x) / n)) * math.Sqrt(sum_y2 - ((sum_y * sum_y) / n))

	if denominator == 0 {
		return 0
	} else {
		return (sum_xy - ((sum_x * sum_y) / n))/ denominator
	}

	return 0;
}

func distance(r1 Rating, r2 Rating, r float64) float64 {
	var distance float64
	var xyr float64
	for k, v := range r1.movies {
		for k2, v2 := range r2.movies {
			if k == k2 && v != 0 && v2 != 0 {
				xyr = math.Pow(math.Abs(v - v2), r)
			}
		}
	}
	//fmt.Println("xyr: ", xyr)
	distance = math.Pow(xyr, 1/ r)
	return distance;
}

func recommend(s string) {
	var distances map[string]Rating = make(map[string]Rating, 1)
	var distancesSlices []Rating = make([]Rating, 1)
	var userArray []string = make([]string, 1)
	for k, v := range users {
		if k != s {
			r := pearson(users[s], users[k])
			fmt.Println("R = ", r)
			v.distance = distance(users[s], users[k], r)
			distances[k] = v
			distancesSlices = append(distancesSlices, v)
			userArray = append(userArray, k)
		}
	}
	quickSort(distancesSlices, 0, len(distancesSlices) - 1)
	fmt.Println("distances = ", distances)
	fmt.Println("user = ", userArray)
}

func quickSort (arr []Rating, low int, high int) {
	if low >= high {
		return
	}
	middle := low + (high - low) / 2

	pivot := arr[middle].distance

	i, j := low, high

	for i < j {
		for arr[i].distance < pivot {
			i++
		}
		for arr[j].distance > pivot {
			j--
		}
		if i <= j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--
		}
	}

	if low < j {
		quickSort(arr, low, j)
	}

	if high > i {
		quickSort(arr, i, high)
	}

}

func findNumber(arr []float64, low int, high int, value float64) {

	if low > high {
		return
	}
	middle := low + (high - low) / 2
	fmt.Println("middle: %v, value at middle: %v", middle, arr[middle])
	if arr[middle] == value {
		fmt.Println("found: %v at %v", value, middle)
		return
	}
	findNumber(arr, low, middle -1, value)
	findNumber(arr, middle + 1, high, value)
}

func prettyPrinter() {
	fmt.Println("======================>>")
	for k, v := range users {
		fmt.Println(k)
		for k2, v2 := range v.movies {
			fmt.Println("	", k2, ":", v2)
		}
	}
	fmt.Println("<<======================")
}

func main() {
	initDataForUsers()
	prettyPrinter()
	//fmt.Println(pearson(users["Angelica"], users["Bill"]))
	//fmt.Println(distance(users["Angelica"], users["Bill"], pearson(users["Angelica"], users["Bill"])))
	recommend("Dan")
}
