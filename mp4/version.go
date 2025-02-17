package mp4

import (
	"fmt"
	"strconv"
	"time"
)

var (
	commitVersion string = "v0.30.1"    // May be updated using build flags
	commitDate    string = "1667671435" // commitDate in Epoch seconds (may be overridden using build flags)
)

// GetVersion - get version and also commitHash and commitDate if inserted via Makefile
func GetVersion() string {
	seconds, _ := strconv.Atoi(commitDate)
	if commitDate != "" {
		t := time.Unix(int64(seconds), 0)
		return fmt.Sprintf("%s, date: %s", commitVersion, t.Format("2006-01-02"))
	}
	return commitVersion
}
