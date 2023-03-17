package utils

import (
	"fmt"
	"log"
	"math"
	"receipt-processor-challenge/src/entities"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Note: Added extra logs for better readability, can be reduced as needed later
func InTimeSpan(check time.Time) bool {
	// Assuming time is in 24 hour format i.e, 02:00pm is 14:00
	start, _ := time.Parse("15:04", "14:01")
	end, _ := time.Parse("15:04", "15:59")

	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func calculatePoints_totalPrice(total string) int {
	points_based_on_total := 0
	total_in_float, error := strconv.ParseFloat(total, 64)
	if error != nil {
		fmt.Println(error)
	} else {
		if math.Round(total_in_float) == total_in_float {
			points_based_on_total += 50
			log.Println("Total price", total_in_float, "is a  total is a round dollar amount with no cents so incrementing points by 50")
		}
	}

	if math.Mod(total_in_float, 0.25) == 0 {
		points_based_on_total += 25
		log.Println("Incrementing points by 25 since total", total_in_float, "is a mutliple of 0.25")
	}

	return points_based_on_total
}

func calculatePoints_processItems(items []entities.Item) int {
	points_based_on_items := 0
	for _, item := range items {

		if len(strings.Trim(item.ShortDescription, " "))%3 == 0 {
			price, error := strconv.ParseFloat(item.Price, 64)
			if error != nil {
				fmt.Println(error)
			} else {
				log.Println("Price", price, "is rounded to", math.Round(price), "and multiplied by .2 =", math.Round(price)*0.2)
				points_based_on_items += int(math.Round(price * 0.2))
				log.Println("Trimmed length of item description is", len(strings.Trim(item.ShortDescription, " ")), "and divisibility by 3 is:", len(strings.Trim(item.ShortDescription, " "))%3 == 0)
			}
		}
	}
	return points_based_on_items
}

func calculatePoints_oddDays(purchaseDate string) int {
	points_if_day_is_odd := 0
	day, error := strconv.ParseInt(strings.Split(purchaseDate, "-")[1], 36, 36)
	if error != nil {
		fmt.Println(error)
	} else {
		if day%2 != 0 {
			points_if_day_is_odd += 6
			log.Println("Day", day, "is odd, incrementing 6 points")
		}
	}
	return points_if_day_is_odd
}

func calculatePoints_timeWithinDuration(purchaseTime string) int {
	points_if_time_lies_within_duration := 0
	time, error := time.Parse("15:04", purchaseTime)
	if error != nil {
		fmt.Println(error)
	} else {
		if InTimeSpan(time) {
			points_if_time_lies_within_duration += 10
			log.Println("Time", time, "lies within 2-4 pm, adding 10 points")
		}
	}
	return points_if_time_lies_within_duration
}

func calculatePoints_alphanumericString(retailer string) int {
	alphanumeric_points := 0
	for _, c := range retailer {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			alphanumeric_points += 1
		}
	}
	log.Println("Retailer name consists of a total of", alphanumeric_points, "letters and numbers")
	return alphanumeric_points
}

func Calculate(receipt entities.Receipt) int {
	points := 0

	points += calculatePoints_alphanumericString(receipt.Retailer)

	points += calculatePoints_totalPrice(receipt.Total)

	points += (int(len(receipt.Items) / 2)) * 5
	log.Println("Incrementing 5 points for every 2 items since there's", len(receipt.Items), "item(s)")

	points += calculatePoints_processItems(receipt.Items)

	points += calculatePoints_oddDays(receipt.PurchaseDate)

	points += calculatePoints_timeWithinDuration(receipt.PurchaseTime)
	log.Println("Final points are", points)

	return points
}
