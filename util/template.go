package util

import "os"

func CreateTemplate(name string) {
	os.Mkdir(name, 0777)

	textFile := "./" + name + "/input.txt"
	if _, err := os.Stat(textFile); err != nil {
		os.WriteFile(textFile, []byte(""), 0777)
	}

	filename := "./" + name + "/" + name + ".go"
	if _, err := os.Stat(filename); err == nil {
		// Return early if the file already exists
		return
	}

	contents := "package " + name + "\n"
	contents += "\n"
	contents += "import \"github.com/bill-kerr/advent-of-code-2023/util\"\n"
	contents += "\n"
	contents += "func part1(lines []string) {\n"
	contents += "\t\n"
	contents += "}\n"
	contents += "\n"
	contents += "func part2(lines []string) {\n"
	contents += "\t\n"
	contents += "}\n"
	contents += "\n"
	contents += "func Run() {\n"
	contents += "\tlines := util.OpenAndRead(\"./" + name + "/input.txt\")\n"
	contents += "\n"
	contents += "\tpart1(lines)\n"
	contents += "\tpart2(lines)\n"
	contents += "}\n"

	os.WriteFile(filename, []byte(contents), 0777)
}
