package lang

import (
	"github.com/Xuanwo/go-locale"
	"github.com/gogf/gf/text/gstr"
	"sync"
)

type Lang struct {
	currentLangMap map[string]string
}

var lang *Lang
var langOnce sync.Once

// Instance 单例
func Instance() *Lang {
	langOnce.Do(func() {
		lang = &Lang{}
		lang.init()
	})
	return lang
}

func (l *Lang) init() {
	l.currentLangMap = make(map[string]string)
	lan, err := locale.Detect()
	if err == nil {
		if gstr.Pos(lan.String(), "zh") == 0 {
			l.currentLangMap = zh
		}
	}
}

func T(str string) string {
	if re, ok := Instance().currentLangMap[str]; ok {
		return re
	}
	return str
}
