package helper

import (
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

//GoPaths returns the GOPATH as an array of paths
func GoPaths() []string {
	return strings.Split(os.Getenv("GOPATH"), ":")
}

// ExitOnError logs error message in fatal mode.
func ExitOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s\n%s", msg, err.Error())
	}
}

// PanicOnError logs error message in fatal mode.
func PanicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

//IsWindows detect os windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

//SortedKeys return slice of sorted map keys
func SortedKeys(theMap map[string]interface{}) []string {
	keys := make([]string, len(theMap))
	i := 0
	for k := range theMap {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}

//ValueOrDefault ...
func ValueOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

//ActualYear return the actual year as string
func ActualYear() string {
	return strconv.Itoa(time.Now().Year())
}
