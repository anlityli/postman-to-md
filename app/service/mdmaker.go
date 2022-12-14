package service

var MdMaker = &mdMakerService{}

type mdMakerService struct {
}

func (s *mdMakerService) Title(data string, level int) (re string) {
	if level <= 0 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	for i := 0; i < level; i++ {
		re += "#"
	}
	return re + " " + data + "\n"
}

func (s *mdMakerService) Text(data string) (re string) {
	return data + "   \n"
}

func (s *mdMakerService) Table(header []string, data [][]string) (re string) {
	re += "\n"
	re += s.tableRow(header)
	re += s.tableRow(header, "-----")

	for _, row := range data {
		re += s.tableRow(row)
	}
	return re + "\n"
}

func (s *mdMakerService) tableRow(data []string, fillStr ...string) (re string) {
	for i, item := range data {
		if len(fillStr) > 0 {
			item = fillStr[0]
		}
		if i == 0 {
			re += "| " + item
		} else if i == len(data)-1 {
			re += " | " + item + " |\n"
		} else {
			re += " | " + item
		}
	}
	return re
}

func (s *mdMakerService) Code(data string) (re string) {
	re += "```\n"
	re += data + "\n"
	re += "```"
	return re + "\n"
}

func (s *mdMakerService) Quote(data string) (re string) {
	return "> " + data + "\n"
}
