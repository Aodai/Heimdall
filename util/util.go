package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ConvertToSeconds takes an uptime string in the format received by the game server and converts it to seconds.
func ConvertToSeconds(uptime string) int {
	days := parseDays(&uptime)
	hours := parseHours(&uptime)
	minutes := parseMinutes(&uptime)
	seconds := parseSeconds(&uptime)
	up := days*86400 + hours*3600 + minutes*60 + seconds
	return up
}

// TODO: Don't forget to clean this up and create a generic method to specify what to parse.
func parseDays(uptime *string) int {
	re, err := regexp.Compile("[0-9]+( Day)")
	if err != nil {
		fmt.Println(err.Error())
	}
	days := re.FindString(*uptime)
	days = strings.Replace(days, " Day", "", -1)
	day, err := strconv.Atoi(days)
	if err != nil {
		fmt.Println(err.Error())
	}
	return day
}

func parseHours(uptime *string) int {
	re, err := regexp.Compile("[0-9]+( Hour)")
	if err != nil {
		fmt.Println(err.Error())
	}
	hours := re.FindString(*uptime)
	hours = strings.Replace(hours, " Hour", "", -1)
	hour, err := strconv.Atoi(hours)
	if err != nil {
		fmt.Println(err.Error())
	}
	return hour
}

func parseMinutes(uptime *string) int {
	re, err := regexp.Compile("[0-9]+( Minute)")
	if err != nil {
		fmt.Println(err.Error())
	}
	minutes := re.FindString(*uptime)
	minutes = strings.Replace(minutes, " Minute", "", -1)
	minute, err := strconv.Atoi(minutes)
	if err != nil {
		fmt.Println(err.Error())
	}
	return minute
}

func parseSeconds(uptime *string) int {
	re, err := regexp.Compile("[0-9]+( Second)")
	if err != nil {
		fmt.Println(err.Error())
	}
	seconds := re.FindString(*uptime)
	seconds = strings.Replace(seconds, " Second", "", -1)
	second, err := strconv.Atoi(seconds)
	if err != nil {
		fmt.Println(err.Error())
	}
	return second
}
