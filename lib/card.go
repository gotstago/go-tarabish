package passwordless

import (
	//"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Card struct {
	Suit       string
	Rank       string
	SuitSymbol string
	RankSymbol string
}

// Card
//type Card struct {
//	Suit      string
//	Rank       string
//}

//
func (c *Card) Symbol() string {
	return c.Suit[:1] + c.Rank
}

func (c *Card) RankValue() int {
	return sort.SearchStrings(RANKS, c.Rank)
}

func (c *Card) SuitValue() int {
	return sort.SearchStrings(SUITS, c.Suit)
}

//
func (c *Card) ToString() string {
	return c.Symbol() //GLYPH[c.Symbol()]
}

// Deck
type Deck struct {
	Cards []Card
}

//
func (d *Deck) Hand() Hand {
	// pop
	hand_cards := d.Cards[0:DEFAULT_HAND_SIZE]
	// pop
	d.Cards = d.Cards[DEFAULT_HAND_SIZE:]
	return Hand{"sample", hand_cards}
}

//
func (d *Deck) Shuffle() *Deck {
	a := d.Cards
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}

	//	src := d.Cards
	//	rand.Seed(time.Now().UnixNano())
	//	perm := rand.Perm(len(d.Cards))
	//	for i, v := range perm {
	//		d.Cards[v] = src[i]
	//	}
	return d
}

func (d *Deck) ToString() string {
	var cardsSymbol []string
	for _, card := range d.Cards {
		cardsSymbol = append(cardsSymbol, card.Symbol()) //GLYPH[card.Symbol()])
	}
	return strings.Join(cardsSymbol, " ")
}

// El
//type Hand struct {
//	Cards []Card
//}

type Hand struct {
	Name  string
	Cards []Card
}

func (h *Hand) Deal() Card {
	// pop
	card := h.Cards[len(h.Cards)-1:]
	h.Cards = h.Cards[0 : len(h.Cards)-1]
	return card[0]
}

//
func (h *Hand) ToString() string {
	var cardsSymbol []string
	for _, card := range h.Cards {
		cardsSymbol = append(cardsSymbol, card.Symbol()) //GLYPH[card.Symbol()])
	}
	return strings.Join(cardsSymbol, " ")
}

var (
	SUITS             []string          = []string{"club", "diamond", "heart", "spade"}
	RANKS             []string          = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "j", "q", "k"}
	DEFAULT_HAND_SIZE                   = 9 //
	GLYPH             map[string]string = map[string]string{
		"s1": "\U0001f0a1", "h1": "\U0001f0b1", "d1": "\U0001f0c1", "c1": "\U0001f0d1",
		"s2": "\U0001f0a2", "h2": "\U0001f0b2", "d2": "\U0001f0c2", "c2": "\U0001f0d2",
		"s3": "\U0001f0a3", "h3": "\U0001f0b3", "d3": "\U0001f0c3", "c3": "\U0001f0d3",
		"s4": "\U0001f0a4", "h4": "\U0001f0b4", "d4": "\U0001f0c4", "c4": "\U0001f0d4",
		"s5": "\U0001f0a5", "h5": "\U0001f0b5", "d5": "\U0001f0c5", "c5": "\U0001f0d5",
		"s6": "\U0001f0a6", "h6": "\U0001f0b6", "d6": "\U0001f0c6", "c6": "\U0001f0d6",
		"s7": "\U0001f0a7", "h7": "\U0001f0b7", "d7": "\U0001f0c7", "c7": "\U0001f0d7",
		"s8": "\U0001f0a8", "h8": "\U0001f0b8", "d8": "\U0001f0c8", "c8": "\U0001f0d8",
		"s9": "\U0001f0a9", "h9": "\U0001f0b9", "d9": "\U0001f0c9", "c9": "\U0001f0d9",
		"s10": "\U0001f0aa", "h10": "\U0001f0ba", "d10": "\U0001f0ca", "c10": "\U0001f0da",
		"sj": "\U0001f0ab", "hj": "\U0001f0bb", "dj": "\U0001f0cb", "cj": "\U0001f0db",
		"sq": "\U0001f0ad", "hq": "\U0001f0bd", "dq": "\U0001f0cd", "cq": "\U0001f0dd",
		"sk": "\U0001f0ae", "hk": "\U0001f0be", "dk": "\U0001f0ce", "ck": "\U0001f0de",
	}
)
