package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"postman-to-md/app/model"
)

var Console = &consoleService{}

type consoleService struct {
}

func (s *consoleService) Run() {
	parser, err := gcmd.Parse(g.MapStrBool{
		"h,help":   false,
		"i,input":  true,
		"o,output": true,
		"s,split":  false,
	})
	if err != nil {
		panic(err)
	}
	allOpt := parser.GetOptAll()
	// help
	_, hOk := allOpt["h"]
	_, helpOk := allOpt["help"]
	if len(allOpt) <= 0 || hOk || helpOk {
		Help.Run(parser.GetArg(0))
		return
	}

	// split option
	_, sOk := allOpt["s"]
	_, splitOk := allOpt["split"]
	split := false
	if sOk || splitOk {
		split = true
	}

	// input file path
	inputPath := parser.GetOpt("i")
	// output file path
	outputPath := parser.GetOpt("o")

	param := &model.RunParam{
		InputPath:  inputPath,
		OutputPath: outputPath,
		Split:      split,
	}
	err = PmReader.Run(param)
	if err != nil {
		panic(err)
	}
}
