package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

var (
	inputFilePath     string
	outputFilePath    string
	inputFileContents string

	notesSharp = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	notesFlat  = [12]string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

	noteToIndice = map[string]int{"C": 0, "C#": 1, "Db": 1, "D": 2, "D#": 3, "Eb": 3, "E": 4, "F": 5, "F#": 6, "Gb": 6, "G": 7, "G#": 8, "Ab": 8, "A": 9, "A#": 10, "Bb": 10, "B": 11, "": -1}
	indiceToNote = map[int]string{0: "c", 1: "c+", 2: "d", 3: "e-", 4: "e", 5: "f", 6: "f+", 7: "g", 8: "a-", 9: "a", 10: "b-", 11: "b"}
)

func main() {
	if !getCmdLineArgs() {
		return
	}
	if !readTheFile() {
		color.Red.Println("Error in reading file!")
		return
	}
	if !process() {
		color.Red.Println("Error!")
		return
	}
	color.Cyan.Println("Parsing done!")
	color.Cyan.Println("Play your song using - ")
	color.Green.Println("alda play --file " + outputFilePath)
}

func process() bool {
	// This is the Regular Expression that matches a chord on the ALDAX notation
	// <alda-code-optional>/<chord>(inversion-optional)/<alda-code-optional>
	var re = regexp.MustCompile(`^([^\sCDEFGAB]+)?([CDEFGAB][#b]?)(6|7|M[2467]?|m[67]?|aug|\+|dim7?|o7?|sus[24])?(\([1-4]\))?(\/\S+)?$`)
	// This contains the output that must be written to the file
	var outputFileContents string
	// splitting the file contents by new line
	var lines = strings.Split(inputFileContents, "\n")
	// iterating through all the lines
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// splitting each line by spaces
		var words = strings.Split(line, " ")
		// iterating through all the words of each line
		for _, word := range words {
			// This will find if the word is a chord
			var match = re.FindStringSubmatch(word)
			if match != nil {
				// Yes. The word is a chord from the ALDAX syntax
				// W will convert the chord to an equivalent ALDA syntax
				var success, chord = getAldaChord(match[1], match[2], match[3], match[4], match[5])
				if !success {
					return false
				}
				// We will add that to the output string
				outputFileContents += chord + " "
			} else {
				// The word is not an ALDAX chord. We simply add it to the output string
				outputFileContents += word + " "
			}
		}
		// adding a new line to the output string
		outputFileContents += "\n"
	}
	ioutil.WriteFile(outputFilePath, []byte(outputFileContents), os.ModePerm)
	return true
}

func getAldaChord(extraAldaStuff1, tonic, chordType, inversionStr, extraAldaStuff2 string) (bool, string) {
	extraAldaStuff1 = strings.TrimSuffix(extraAldaStuff1, "/")
	inversionStr = strings.TrimPrefix(inversionStr, "(")
	inversionStr = strings.TrimSuffix(inversionStr, ")")
	inversion, _ := strconv.Atoi(inversionStr)
	inversion = inversion % 4
	extraAldaStuff2 = strings.TrimPrefix(extraAldaStuff2, "/")

	var chord string = ""
	var success bool
	var formula []int
	success, formula = getChordFormula(chordType)
	if !success || formula == nil {
		return false, ""
	}

	var notes []string
	success, notes = getNotes(noteToIndice[tonic], formula, inversion, extraAldaStuff1, extraAldaStuff2)
	if !success || notes == nil {
		return false, ""
	}
	for _, note := range notes {
		chord += note + "/"
	}
	chord = strings.TrimSuffix(chord, "/")

	return true, chord
}

func getChordFormula(chordType string) (bool, []int) {
	switch chordType {
	case "", "M":
		// major
		return true, []int{0, 4, 7}
	case "m":
		// minor
		return true, []int{0, 3, 7}
	case "7":
		// Dominant 7th
		return true, []int{0, 4, 7, 10}
	case "M7":
		// Major 7th
		return true, []int{0, 4, 7, 11}
	case "m7":
		// minor 7th
		return true, []int{0, 3, 7, 10}
	case "aug", "+":
		// augmented
		return true, []int{0, 4, 8}
	case "dim", "o":
		// diminished
		return true, []int{0, 3, 6}
	case "dim7", "o7":
		// diminished 7th
		return true, []int{0, 3, 6, 9}
	case "2", "M2", "sus2":
		// suspended 2
		return true, []int{0, 2, 7}
	case "4", "M4", "sus4":
		// suspended 4
		return true, []int{0, 5, 7}
	case "6", "M6":
		// Major 6th
		return true, []int{0, 4, 7, 9}
	case "m6":
		// Minor 6th
		return true, []int{0, 3, 7, 9}
	default:
		return false, nil
	}
}

func getNotes(tonicIndex int, formula []int, inversion int, extraAldaStuff1 string, extraAldaStuff2 string) (bool, []string) {
	var notes []string
	notesIndex := []int{}

	// computing and adding the the notes with the tonic and the formula
	for _, addition := range formula {
		var noteIndex = tonicIndex + addition
		notesIndex = append(notesIndex, noteIndex)
	}

	for i := 0; i < inversion; i++ {
		tmp := notesIndex[0]
		notesIndex = notesIndex[1:]
		notesIndex = append(notesIndex, tmp)
	}

	if notesIndex[0] > 11 {
		notesIndex[0] = notesIndex[0] - 12
	}
	for i := 1; i < len(notesIndex); i++ {
		for notesIndex[i] < notesIndex[i-1] {
			notesIndex[i] = notesIndex[i] + 12
		}
	}
	for _, noteIndex := range notesIndex {
		if noteIndex < 0 {
			notes = append(notes, extraAldaStuff1+"<"+indiceToNote[noteIndex+12]+extraAldaStuff2+">")
		} else if noteIndex > 11 {
			notes = append(notes, extraAldaStuff1+">"+indiceToNote[noteIndex-12]+extraAldaStuff2+"<")
		} else {
			notes = append(notes, extraAldaStuff1+indiceToNote[noteIndex]+extraAldaStuff2)
		}
	}

	return true, notes
}

// MISC Helper functions

func getCmdLineArgs() bool {
	var cmdLineArgs []string = os.Args[1:]
	var noOfCmdLineArgs int = len(cmdLineArgs)
	if noOfCmdLineArgs == 1 && cmdLineArgs[0] == "--help" {
		fmt.Println("")
		color.Yellow.Println("*SYNTAX*")
		color.Yellow.Println("* for all the chords *")
		color.Green.Println("<alda-code-optional>/<chord>(inversion-optional)/<alda-code-optional>")
		fmt.Println("")
		color.Cyan.Println("-> Alda code and inversion are optional")
		color.Cyan.Println("-> <chord> contains a capital letter denoting the chord tonic, followed by #, b, 7, etc. to specify the actual chord")
		color.Cyan.Println("-> For each chord, you can optionally specify an inversion in brackets (1-4)")
		color.Cyan.Println("-> The 1st <alda-code-optional> is prefixed to every note in the chord and the 2nd is suffixed to every note in the chord. You may use the 2nd <alda-code-optional> to specify the duration of every note played")
		color.Cyan.Println("-> It is good to note that pure Alda provides you with more flexibility. Aldax provides you a bit of abstraction when it comes to chord notation")
		color.Cyan.Println("Therefore, you can use all the features of Alda, but with a sugar coated abstraction for Chords!")
		fmt.Println("")
		color.LightGreen.Println("The chords supported are major, minor, dominant 7th, major 7th, minor 7th, diminished, diminished 7th, augmented, suspended 2 and suspended 4")
		color.Gray.Println("chord examples...")
		color.LightGreen.Println("Major        : CM, D#")
		color.LightGreen.Println("Minor        : Am, Bbm")
		color.LightGreen.Println("Dominant 7th : G7, D7")
		color.LightGreen.Println("Major 7th    : AM7, AbM7")
		color.LightGreen.Println("Minor 7th    : Am7, Abm7")
		color.LightGreen.Println("Dimished     : Gdim, Do")
		color.LightGreen.Println("Dimished 7th : Edim7, Fo7")
		color.LightGreen.Println("Augmented    : Caug, B+")
		color.LightGreen.Println("Suspended 2  : Fsus2, A#sus2, GM2")
		color.LightGreen.Println("Suspended 4  : Fsus4, A#sus4, GM4")
		color.LightGreen.Println("Major 6th    : G6, DM6, A#6")
		color.LightGreen.Println("Minor 6th    : Gm6, Dm6, A#m6")
		fmt.Println()
		color.Gray.Println("examples...")
		color.Green.Println("Csus2(2)/1        : This contains the C suspended 2 chord in the second inversion, for a duration of 1 (Alda notation)")
		color.Green.Println("C Am F G7         : This plays the classic <I vi IV V7> progression with C as tonic, at the default Alda duration")
		color.Green.Println("Checkout https://github.com/haricane8133/aldax/blob/master/examples for more examples. (PRs for more examples welcome)")
		fmt.Println()
		color.Yellow.Println("*USAGE*")
		color.Cyan.Println("1. Write your Alda song in a file, with Aldax chord syntax")
		color.Cyan.Println("2. Use the Aldax binary to parse your file into a pure Alda file")
		color.Cyan.Println("3. Play the generated file using Alda")
		fmt.Println()
		color.Cyan.Println("To call the parser...")
		color.Blue.Println("aldax <inputfilepath> <outputfilepath>")
		color.Cyan.Println("example: aldax input.aldax output.alda && alda play --file output.alda")
		fmt.Println()

		return false
	}
	if noOfCmdLineArgs != 2 {
		fmt.Println()
		color.Yellow.Println("********** ALDAX **********")
		color.Cyan.Println("Alda's chord syntax parser")
		color.Cyan.Println("This binary converts Aldax to Alda; helps you write chords easily")
		fmt.Println()
		color.Green.Println("You need to pass two arguments")
		fmt.Println("Input file path, ")
		fmt.Println("Output file path")
		fmt.Println()
		color.Red.Println("The repository is at https://www.github.com/haricane8133/aldax. For the help, give --help")
		fmt.Println()
		color.Gray.Println("Input file contains the Aldax syntax")
		color.Gray.Println("Output file contains the equivalent Alda code")
		fmt.Println()
		color.Gray.Println("The convention is for input file to have .aldax extension and output file to have .alda extension")
		fmt.Println()
		return false
	}

	inputFilePath = cmdLineArgs[0]
	outputFilePath = cmdLineArgs[1]
	return true
}

func readTheFile() bool {
	dat, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	inputFileContents = string(dat)
	return true
}
