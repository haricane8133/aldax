# Aldax

Aldax is a chord parser for the musical language Alda


## Installation

1. Install [Go](https://golang.org/dl/)
2. Download the repo
3. Go to the project root 

```bash
go install
# This installs the Binary in the default $GOPATH/bin directory. If it is not in the path, add it
    # If you use a Bash env, add this line to your bash profile (or zsh likewise)
    # export PATH=$PATH:$(go env GOPATH)/bin
```

For people who do not want to go through the above hassle, download the releases from the releases tab


## Usage

Aldax takes 2 parameters,
1. Path of the Aldax code file for input
2. Path of an Alda code file for output

The input file contains Alda code and can contain Aldax based chord syntax. The chords will be parsed into pure Alda syntax by the program, which can then be played by Alda

```bash
aldax input-file-path output-file-path
alda play --file output-file-path
```

## Videos

[What is Aldax?](https://youtu.be/6x3iGFwMmuk)
Tells you a bit about what is Aldax and why you need it...

[How to use Aldax?](https://youtu.be/F9mE1id8ixM)
It covers Installation and Usage from above

[Writing a song with Aldax](https://youtu.be/nvGb9fhiIFg)
Tune in to watch how I write the intro of the song "Let it be" by The Beatles, in Alda, using the Chords syntax of Aldax

## Aldax chord Syntax

Each chord within the Aldax code file, is represented in the following format

alda-code/chord(inversion)/alda-code
* alda-code and inversion are optional. It is good to note that pure Alda provides you with more flexibility
* Aldax provides you with an easy abstraction for chords
* "chord" contains a capital letter denoting the chord tonic, followed by #, b, 7, etc. to specify the actual chord. Have a look at the chords supported for knowing what to use
* For each chord, you can optionally specify an "inversion" in brackets (1-4)
* The 1st <alda-code-optional> is prefixed to every note in the chord and the 2nd is suffixed to every note in the chord. You may use the 2nd <alda-code-optional> to specify the duration of every note played

Note: YOU NEED TO USE CAPITALS FOR THE "CHORD". SEE THE EXAMPLES TO KNOW HOW

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
* Major 6th    : G6, DM6, A#6
* Minor 6th    : Gm6, Dm6, A#m6

### Full Syntax - aldax file contents

#### Csus2(2)/1/
This contains the C suspended 2 chord in the second inversion, for a duration of 1 (Alda notation)

#### C Am F G7
This plays the classic < I vi IV V7> progression with C as tonic, at the default Alda duration


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[EPL-v1.0](https://www.eclipse.org/legal/epl-v10.html)