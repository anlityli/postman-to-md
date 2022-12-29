package service

import "github.com/gogf/gf/util/gconv"

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

func (s *mdMakerService) RowCode(data string) (re string) {
	return "```" + data + "```"
}

func (s *mdMakerService) Quote(data string) (re string) {
	return "> " + data + "\n"
}

func (s *mdMakerService) Color(data string, color string) (re string) {
	return "<font color=" + color + ">" + data + "</font>"
}

func (s *mdMakerService) Bold(data string) (re string) {
	return "**" + data + "**"
}

func (s *mdMakerService) Italic(data string) (re string) {
	return "*" + data + "*"
}

func (s *mdMakerService) BoldAndItalic(data string) (re string) {
	return "***" + data + "***"
}

func (s *mdMakerService) Lists(data []string, order ...bool) (re string) {
	re = ""
	for key, item := range data {
		if len(order) > 0 && order[0] {
			re += gconv.String(key+1) + ". " + item + "\n"
		} else {
			re += "- " + item + "\n"
		}
	}
	return re
}

func (s *mdMakerService) HorizontalRule() (re string) {
	return "\n---\n\n"
}

func (s *mdMakerService) Link(title, link string) (re string) {
	return "[" + title + "](" + link + ")"
}

func (s *mdMakerService) Url(data string) (re string) {
	return "<" + data + ">"
}

func (s *mdMakerService) Image(imagePath, altText string, link ...string) (re string) {
	if len(link) > 0 && link[0] != "" {
		return "[![" + altText + "](" + imagePath + ")](" + link[0] + ")"
	} else {
		return "![" + altText + "](" + imagePath + ")"
	}
}
