# Aldax

Aldax is a chord parser for the musical language Alda


## Installation

1. Install [Go](https://golang.org/dl/)
2. Download the repo
3. Go to the project root 

```bash
go install
```

For people who do not want to go through the above hassle, download the releases from the releases tab


## Usage

Aldax takes 2 parameters,
1. Path of the Aldax code file for input
2. Path of an Alda code file for output

The input file contains Alda code and can contain Aldax based chord syntax. The chords will be parsed into pure Alda syntax by the program, which can then be played by Alda

```bash
aldax < input-file-path> < output-file-path>
alda play --file < output-file-path>
```


## Aldax chord Syntax

Each chord within the Aldax code file, is represented in the following format

< other-alda-code>/< chord>/< chord-base-note>/< duration>/< other-alda-code>

* < other-alda-code> denotes Alda code that does not indicate a chord, that you may want to chain up with the chord
* < chord> denotes a chord
* < chord-base-note> is used for chord inversions
* < duration> is the Alda based duration for the whole chord

Note: YOU NEED TO USE CAPITALS FOR THE "CHORD" AND THE "CHORD-BASE-NOTE". SEE THE EXAMPLES TO KNOW HOW

"Chord base note", "duration", and "other alda code" are optional. They provide extra features to all. It is good to note that pure Alda provides you with more flexibility

The chord types supported here are major, minor, dominant 7th, major 7th, minor 7th, diminished, diminished 7th, augmented, suspended 2 and suspended 4. If you want more chords, then you are probably an expert who can tackle Alda directly XD


## Examples

### Chords
* Major        : CM, D#
* Minor        : Am, Bbm
* Dominant 7th : G7, D7
* Major 7th    : AM7, AbM7
* Minor 7th    : Am7, Abm7
* Dimished     : Gdim, Do
* Dimished 7th : Edim7, Fo7
* Augmented    : Caug, B+
* Suspended 2  : Fsus2, A#sus2, GM2
* Suspended 4  : Fsus4, A#sus4, GM4

### Full Syntax - aldax file contents

#### Csus2/D/1/<d>
This contains the major chord, with D as the base note, for a duration of 1 (Alda notation) and also has some Alda code at the end (< d>)

#### C Am F G7
This plays the classic < I vi IV V7> progression with C as tonic, at the default Alda duration


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[EPL-v1.0](https://www.eclipse.org/legal/epl-v10.html)