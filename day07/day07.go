package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bill-kerr/advent-of-code-2023/util"
)

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var cardValues map[rune]int = map[rune]int{}

var cardValuesPart1 map[rune]int = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

var cardValuesPart2 map[rune]int = map[rune]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type hand struct {
	cards    []rune
	bid      int
	handType handType
	highCard rune
	counts   map[rune]int
}

func newHand(handString string, bid int) hand {
	hand := hand{
		cards:  make([]rune, len(handString)),
		bid:    bid,
		counts: make(map[rune]int),
	}

	pairs := 0
	triples := 0

	for i, card := range handString {
		hand.cards[i] = card
		hand.counts[card] += 1

		if cardValues[card] > cardValues[hand.highCard] {
			hand.highCard = card
		}

		if hand.counts[card] == 5 {
			hand.handType = fiveOfAKind
			continue
		}

		if hand.counts[card] == 4 {
			hand.handType = fourOfAKind
			continue
		}

		if hand.counts[card] == 4 {
			hand.handType = fourOfAKind
			continue
		}

		if hand.counts[card] == 3 && pairs == 2 {
			hand.handType = fullHouse
		}

		if hand.counts[card] == 3 && pairs == 1 {
			hand.handType = threeOfAKind
			triples++
			continue
		}

		if hand.counts[card] == 2 && triples == 1 {
			hand.handType = fullHouse
			pairs++
			continue
		}

		if hand.counts[card] == 2 && pairs == 1 {
			hand.handType = twoPair
			pairs++
			continue
		}

		if hand.counts[card] == 2 && pairs == 0 {
			hand.handType = onePair
			pairs++
			continue
		}
	}

	return hand
}

func (h *hand) less(otherHand hand) bool {
	if h.handType > otherHand.handType {
		return false
	}

	if h.handType < otherHand.handType {
		return true
	}

	for i := 0; i < len(h.cards); i++ {
		if cardValues[h.cards[i]] > cardValues[otherHand.cards[i]] {
			return false
		}

		if cardValues[h.cards[i]] < cardValues[otherHand.cards[i]] {
			return true
		}
	}

	panic("hands cannot be the same")
}

func (h *hand) applyJokers() {
	jokerCount := h.counts['J']
	if jokerCount == 0 {
		return
	}

	if h.handType == fourOfAKind && jokerCount >= 1 {
		h.handType = fiveOfAKind
	} else if (h.handType == fullHouse || h.handType == threeOfAKind) && jokerCount >= 2 {
		h.handType = fiveOfAKind
	} else if h.handType == threeOfAKind && jokerCount == 1 {
		h.handType = fourOfAKind
	} else if h.handType == twoPair && jokerCount == 2 {
		h.handType = fourOfAKind
	} else if h.handType == twoPair && jokerCount == 1 {
		h.handType = fullHouse
	} else if h.handType == onePair && (jokerCount == 1 || jokerCount == 2) {
		h.handType = threeOfAKind
	} else if h.handType == highCard && jokerCount == 1 {
		h.handType = onePair
	}
}

func part1(lines []string) {
	cardValues = cardValuesPart1
	hands := make([]hand, len(lines))

	for i, line := range lines {
		handString := line[:5]
		bid, _ := strconv.ParseInt(strings.Split(line, " ")[1], 10, 64)
		hands[i] = newHand(handString, int(bid))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].less(hands[j])
	})

	winnings := 0
	for i, hand := range hands {
		winnings += hand.bid * (i + 1)
	}

	fmt.Println(winnings)
}

func part2(lines []string) {
	cardValues = cardValuesPart2
	hands := make([]hand, len(lines))

	for i, line := range lines {
		handString := line[:5]
		bid, _ := strconv.ParseInt(strings.Split(line, " ")[1], 10, 64)
		hands[i] = newHand(handString, int(bid))
		hands[i].applyJokers()
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].less(hands[j])
	})

	winnings := 0
	for i, hand := range hands {
		winnings += hand.bid * (i + 1)
	}

	fmt.Println(winnings)
}

func Run() {
	lines := util.OpenAndRead("./day07/input.txt")

	part1(lines)
	part2(lines)
}
