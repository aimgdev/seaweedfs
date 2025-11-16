package version

import (
	"fmt"

	"github.com/seaweedfs/seaweedfs/weed/util"
)

var (
	MAJOR_VERSION  = int32(4)
	MINOR_VERSION  = int32(00)
	VERSION_NUMBER = fmt.Sprintf("%d.%02d", MAJOR_VERSION, MINOR_VERSION)
	VERSION        = util.SizeLimit + " " + VERSION_NUMBER
	COMMIT         = ""
	CUSTOM_BUILD   = "jet r4"
)

func Version() string {
	if CUSTOM_BUILD != "" {
		return VERSION + " " + COMMIT + " [" + CUSTOM_BUILD + "]"
	}
	return VERSION + " " + COMMIT
}
