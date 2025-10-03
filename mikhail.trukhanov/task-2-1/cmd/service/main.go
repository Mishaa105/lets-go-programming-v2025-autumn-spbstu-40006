package main

import "fmt"

func main() {

	var (
		num_of_departments, num_of_workers, temp int
		sign                                     string
	)
	if _, err := fmt.Scan(&num_of_departments); err != nil {
		fmt.Println("Invalid input")
		return
	}

	for i := 0; i != num_of_departments; i++ {
		min_temp, max_temp := 15, 30
		if _, err := fmt.Scan(&num_of_workers); err != nil {
			fmt.Println("Invalid input")
			return
		}
		for j := 0; j != num_of_workers; j++ {
			if _, err := fmt.Scan(&sign, &temp); err != nil {
				fmt.Println("Invalid input")
				return
			}
			if sign == ">=" && temp > min_temp {
				min_temp = temp
			}
			if sign == "<=" && temp < max_temp {
				max_temp = temp
			}
			if max_temp >= min_temp {
				fmt.Println(min_temp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
