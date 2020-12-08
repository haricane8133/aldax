# Aldax

Aldax is a chord parser for the musical language [Alda](https://alda.io/)
Chords are given easier syntax; Aldax converts into Alda digestable syntax

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

NOTE: You need to [install Alda](https://alda.io/install/) before proceeding.
(After Aldax converts your musical score into Alda's format, you need to use Alda to play)

## Usage

1. Write a musical score using Alda's syntax augmented with Aldax's chord syntax (given below)
2. Run the musical score through Aldax (instructions given below)
3. Give the output file to Alda to play your score (instructions given below)

### Aldax Instructions
Aldax takes 2 parameters,
1. Path of the Aldax code file for input
2. Path of an Alda code file for output

The input file contains Alda code and can contain Aldax based chord syntax. The chords will be parsed into pure Alda syntax by the program, which can then be played by Alda

```bash
# create Alda compatible score from your Aldax score
aldax input-file-path output-file-path

# play the converted Alda score
alda play --file output-file-path
```

## Videos

[What is Aldax?](https://youtu.be/6x3iGFwMmuk)
Tells you a bit about what is Aldax and why you need it...

[How to use Aldax?](https://youtu.be/F9mE1id8ixM)
It covers Installation and Usage from above

[Writing a song with Aldax](https://youtu.be/nvGb9fhiIFg)
Tune in to watch how I write the intro of the song "Let it be" by The Beatles, in Alda, using the Chords syntax of Aldax

To learn how to write songs in Alda, have a look at the wonderful article series by the Author of Alda, at https://blog.djy.io/writing-music-with-alda-1/

## Aldax chord Syntax

In your score, for chords alone, you can use Aldax's easy syntax. You are free to use the rest of Alda features in your score too

Each chord within the Aldax score, is represented in the following format

alda-code/chord(inversion)/alda-code

* alda-code and inversion are optional
* "chord" contains a capital letter denoting the chord tonic, followed by #, b, 7, etc. to specify the actual chord. Have a look at the Chords Notation below for knowing what to use
* For each chord, you can optionally specify an "inversion" in brackets (1-4), to invert the chord. Aldax keeps the resultant chord's root at the current Octave and the rest of the notes follow after it
* The 1st <alda-code> is prefixed to every note in the chord and the 2nd one is suffixed to every note in the chord. For instance, you may use the 2nd <alda-code> to specify the duration of every note played. (see the song examples in the [/examples](https://github.com/haricane8133/aldax/tree/master/examples) folder)

Note
1. You need a space to separate two chords or a chord from some other Alda code
2. YOU NEED TO USE CAPITALS FOR THE "CHORD". SEE THE CHORD NOTATION and VOICINGS BELOW TO KNOW HOW

The chord types supported here are major, minor, dominant 7th, major 7th, minor 7th, diminished, diminished 7th, augmented, suspended 2 and suspended 4. If you want more chords, then you are probably an expert who can tackle Alda directly XD


## Chord Notation
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

## Chord Voicings

#### Csus2(2)/1/
This contains the C suspended 2 chord in the second inversion, for a duration of 1 (Alda notation)

#### C Am F G7
This plays the classic < I vi IV V7> progression with C as tonic, at the default Alda duration

Note: There are song examples in the [/examples](https://github.com/haricane8133/aldax/tree/master/examples) folder



## Contributing
Pull requests are welcome on Code changes, feature additions, song examples and anything else!


## License
[EPL-v1.0](https://www.eclipse.org/legal/epl-v10.html)