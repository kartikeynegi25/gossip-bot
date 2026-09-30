package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var botNames = []string{
	"Kartikey",
	"Rohit",
	"Siddharth",
	"Amit",
	"Ankit",
	"Rahul",
	"Vikram",
	"Nikhil",
	"Arjun",
	"Rakesh",
	"Deepak",
	"Saurabh",
	"Manish",
	"Ravi",
	"Ayush",
	"Aditya",
	"Pranav",
	"Harsh",
	"Aakash",
	"Shubham",
}

type Rumor struct {
	original string

	subject string
	action  string
	details string

	prefix string
	time   string
	drama  string
}

var linkingWords = map[string]bool{
	"is":   true,
	"am":   true,
	"are":  true,
	"was":  true,
	"were": true,
}

func parseRumor(text string) Rumor {
	words := strings.Fields(text)

	if len(words) == 0 {
		return Rumor{
			original: text,
			details:  text,
		}
	}

	for i, word := range words {
		lower := strings.ToLower(word)

		if linkingWords[lower] && i > 0 {
			return Rumor{
				original: text,
				subject:  strings.Join(words[:i], " "),
				action:   lower,
				details:  strings.Join(words[i+1:], " "),
			}
		}
	}


	if len(words) == 1 {
		return Rumor{
			original: text,
			details:  text,
		}
	}

	if len(words) == 2 {
		return Rumor{
			original: text,
			subject:  words[0],
			action:   words[1],
		}
	}

	return Rumor{
		original: text,
		subject:  words[0],
		action:   words[1],
		details:  strings.Join(words[2:], " "),
	}
}


func rumorText(r Rumor) string {
	if r.subject == "" {
		return r.details
	}

	text := r.subject + " " + r.action

	if r.details != "" {
		text += " " + r.details
	}

	if r.time != "" {
		text += " " + r.time
	}

	if r.drama != "" {
		text += " " + r.drama
	}

	if r.prefix != "" {
		text = r.prefix + " " + text
	}

	return text
}

func addPrefix(r Rumor) Rumor {
	if r.prefix != "" {
		return r
	}

	prefixes := []string{
		"Apparently,",
		"I heard",
		"People are saying",
		"Someone told me",
	}

	r.prefix = prefixes[rand.Intn(len(prefixes))]

	return r
}

func addTime(r Rumor) Rumor {
	if r.time != "" {
		return r
	}

	times := []string{
		"yesterday",
		"last night",
		"earlier",
		"this morning",
		"recently",
	}

	r.time = times[rand.Intn(len(times))]

	return r
}

func addDrama(r Rumor) Rumor {
	if r.drama != "" {
		return r
	}

	drama := []string{
		"and didn't tell anyone",
		"and refused to explain",
		"and everyone noticed",
		"and nobody knows why",
		"and people were shocked",
	}

	r.drama = drama[rand.Intn(len(drama))]

	return r
}

func makeDetailsBigger(r Rumor) Rumor {
	if r.details == "" {
		return r
	}

	words := strings.Fields(r.details)

	// We only change adjectives that are already there.
	// This keeps the original meaning mostly intact.
	adjectiveChanges := map[string]string{
		"big":      "huge",
		"small":    "tiny",
		"good":     "amazing",
		"bad":      "terrible",
		"nice":     "amazing",
		"new":      "expensive",
		"old":      "ancient",
		"fast":     "crazy fast",
		"slow":     "really slow",
		"little":   "tiny",
		"large":    "massive",
		"massive":  "enormous",
		"huge":     "enormous",
		"amazing":  "insane",
		"terrible": "awful",
	}

	for i, word := range words {
		lower := strings.ToLower(word)

		if newWord, ok := adjectiveChanges[lower]; ok {
			words[i] = newWord
			r.details = strings.Join(words, " ")
			return r
		}
	}

	// No obvious adjective was found,
	// so leave the rumor alone.
	return r
}

func changeAction(r Rumor) Rumor {

	switch strings.ToLower(r.action) {

	case "ate":
		r.action = "had"

	case "had":
		r.action = "ate"

	case "bought":
		r.action = "got"

	case "got":
		r.action = "bought"

	case "saw":
		r.action = "noticed"

	case "noticed":
		r.action = "saw"

	case "found":
		r.action = "discovered"

	case "discovered":
		r.action = "found"

	case "went":
		r.action = "visited"

	case "visited":
		r.action = "went"

	default:
		// We don't know this verb,
		// so leave it alone.
		return r
	}

	return r
}

func changeQuantity(r Rumor) Rumor {
	words := strings.Fields(r.details)

	for i, word := range words {
		lower := strings.ToLower(word)

		switch lower {
		case "a":
			if i+1 < len(words) {
				words[i] = "two"
				words[i+1] = pluralize(words[i+1])
			}

		case "one":
			if i+1 < len(words) {
				words[i] = "two"
				words[i+1] = pluralize(words[i+1])
			}

		case "two":
			if i+1 < len(words) {
				words[i] = "three"
				words[i+1] = pluralize(words[i+1])
			}

		case "three":
			if i+1 < len(words) {
				words[i] = "five"
				words[i+1] = pluralize(words[i+1])
			}

		default:
			continue
		}

		r.details = strings.Join(words, " ")
		return r
	}

	return r
}

func pluralize(word string) string {
	lower := strings.ToLower(word)

	if strings.HasSuffix(lower, "s") {
		return word
	}

	return word + "s"
}

func mutateRumor(r Rumor, generation int) Rumor {

	// First few people usually pass the rumor
	// along without changing it.
	if generation <= 5 {
		if rand.Intn(5) == 0 {
			return addPrefix(r)
		}

		return r
	}

	// Middle of the gossip chain.
	if generation <= 12 {

		switch rand.Intn(6) {
		case 0:
			return addPrefix(r)

		case 1:
			return addTime(r)

		case 2:
			return makeDetailsBigger(r)

		case 3:
			return changeAction(r)

		case 4:
			return changeQuantity(r)

		default:
			return r
		}
	}

	// The last few people are more likely
	// to exaggerate the story.
	switch rand.Intn(6) {
	case 0:
		return addTime(r)

	case 1:
		return makeDetailsBigger(r)

	case 2:
		return changeAction(r)

	case 3:
		return addDrama(r)

	default:
		return r
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("========================================")
	fmt.Println("          GOSSIP MUTATOR BOT")
	fmt.Println("========================================")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a gossip: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Could not read the gossip.")
		return
	}

	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Please enter something.")
		return
	}

	rumor := parseRumor(input)

	fmt.Println()
	fmt.Println("Starting rumor:")
	fmt.Println(input)
	fmt.Println()
	fmt.Println("----------------------------------------")

	for i := 0; i < len(botNames); i++ {

		time.Sleep(500 * time.Millisecond)

		rumor = mutateRumor(rumor, i+1)

		fmt.Printf(
			"%s whispers: %s\n",
			botNames[i],
			rumorText(rumor),
		)
	}

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("              FINAL RUMOR")
	fmt.Println("----------------------------------------")
	fmt.Println(rumorText(rumor))
	fmt.Println()
	fmt.Println("And that's how gossip spreads 💀")
	fmt.Println("========================================")
}
