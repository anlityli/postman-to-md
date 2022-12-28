package service

import (
	"fmt"
	"postman-to-md/library/lang"
)

var Help = &helpService{}

type helpService struct {
}

func (s *helpService) Run(exeFile string) {
	helpStr := `USAGE
    %s [option]

OPTION
    -f,  --file     %s
    -s,  --split    %s
    -h,  --help     %s

`
	fmt.Printf(helpStr, exeFile, lang.T("Postman json file's path."), lang.T("Split the *.md file into a separate directory."), lang.T("Help"))
}
