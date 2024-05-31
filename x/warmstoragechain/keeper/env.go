package keeper

import (
	"os"
)

var (
	CIDExtraOutPutDir string
)

func init() {
	workdir, err := os.Getwd()
	if err == nil {
		CIDExtraOutPutDir += workdir + "/cid_data_dir"
	}

}
