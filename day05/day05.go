package day05

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

type MapEntry struct {
	destinationOffset int
	sourceRange       Range
}

type Almanac struct {
	chapters []Chapter
}

type Chapter struct {
	name    string
	entries []MapEntry
}

type Range struct {
	start int
	end   int
}

func newAlmanac(lines []string) Almanac {
	almanac := Almanac{}

	for i, line := range lines {
		for _, chapterName := range chapterNames {
			if strings.Contains(line, chapterName) {
				almanac.chapters = append(almanac.chapters, Chapter{
					name:    chapterName,
					entries: parseMap(lines, i+1),
				})
			}
		}
	}

	return almanac
}

func (c *Chapter) lookupBySource(source int) int {
	for _, entry := range c.entries {
		if entry.sourceRange.contains(source) {
			return source + entry.destinationOffset
		}
	}

	return source
}

func (c *Chapter) convertRanges(sourceRanges []Range) []Range {
	converted := []Range{}
	current := 0

	for current < len(sourceRanges) {
		sourceRange := sourceRanges[current]

		for _, entry := range c.entries {
			// source range is fully contained within the entry range
			if sourceRange.start >= entry.sourceRange.start && sourceRange.end <= entry.sourceRange.end {
				converted = append(converted, Range{
					start: sourceRange.start + entry.destinationOffset,
					end:   sourceRange.end + entry.destinationOffset,
				})
				break
			}

			// source range starts within the entry range, but ends outside of it
			if sourceRange.start >= entry.sourceRange.start &&
				sourceRange.start <= entry.sourceRange.end &&
				sourceRange.end > entry.sourceRange.end {
				// append the portion of the source range that overlaps with the entry range
				converted = append(converted, Range{
					start: sourceRange.start + entry.destinationOffset,
					end:   entry.sourceRange.end + entry.destinationOffset,
				})

				// add a new (untransformed) source range that represents the leftover portion
				sourceRanges = append(sourceRanges, Range{
					start: entry.sourceRange.end + 1,
					end:   sourceRange.end,
				})
			}
		}

		current++
	}

	return converted
}

func (r *Range) contains(num int) bool {
	return num >= r.start && num <= r.end
}

func part1(lines []string) {
	seeds := util.ParseInts(strings.Split(lines[0], " ")[1:])
	almanac := newAlmanac(lines)
	lowestLocationNumber := int(math.MaxInt64)

	for _, seed := range seeds {
		output := seed

		for _, chapter := range almanac.chapters {
			output = chapter.lookupBySource(output)
		}

		lowestLocationNumber = min(output, lowestLocationNumber)
	}

	fmt.Println(lowestLocationNumber)
}

var chapterNames []string = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func parseMap(lines []string, startIndex int) []MapEntry {
	givenEntries := []MapEntry{}

	index := startIndex
	for index < len(lines) {
		// If the first character isn't a digit, we know it's not a map entry
		if !isFirstCharacterDigit(lines[index]) {
			break
		}

		values := util.ParseInts(strings.Split(lines[index], " "))

		givenEntries = append(givenEntries, MapEntry{
			destinationOffset: values[0] - values[1],
			sourceRange:       Range{start: values[1], end: values[1] + values[2] - 1},
		})

		index++
	}

	slices.SortFunc(givenEntries, func(a, b MapEntry) int {
		if a.sourceRange.start < b.sourceRange.start {
			return 1
		}
		return -1
	})

	entries := []MapEntry{}

	i := 0
	for i < math.MaxInt64 {
		entry := getNextEntry(givenEntries, i)
		i = entry.sourceRange.end + 1
		entries = append(entries, entry)
	}

	return entries
}

func getNextEntry(entries []MapEntry, source int) MapEntry {
	minStart := math.MaxInt64

	for _, entry := range entries {
		if entry.sourceRange.start == source {
			return entry
		}

		if entry.sourceRange.start > source && entry.sourceRange.start < minStart {
			minStart = entry.sourceRange.start
		}
	}

	// In this case, there was no range with a start higher than source
	if minStart == math.MaxInt64 {
		return MapEntry{
			sourceRange:       Range{start: source, end: math.MaxInt64 - 1},
			destinationOffset: 0,
		}
	}

	return MapEntry{
		sourceRange:       Range{start: source, end: minStart - 1},
		destinationOffset: 0,
	}
}

func isFirstCharacterDigit(line string) bool {
	if len(line) < 1 {
		return false
	}

	_, err := util.RtoDigit(rune(line[0]))
	return err == nil
}

func part2(lines []string) {
	seedNumbers := util.ParseInts(strings.Split(lines[0], " ")[1:])
	almanac := newAlmanac(lines)
	minLocationNumber := math.MaxInt64

	for i := 0; i < len(seedNumbers)/2; i++ {
		seedRange := Range{start: seedNumbers[i*2], end: seedNumbers[i*2] + seedNumbers[i*2+1] - 1}

		currentRanges := []Range{seedRange}
		for _, chapter := range almanac.chapters {
			currentRanges = chapter.convertRanges(currentRanges)
		}

		for _, r := range currentRanges {
			minLocationNumber = min(r.start, minLocationNumber)
		}
	}

	fmt.Println(minLocationNumber)
}

func Run() {
	lines := util.OpenAndRead("./day05/input.txt")

	part1(lines)
	part2(lines)
}
