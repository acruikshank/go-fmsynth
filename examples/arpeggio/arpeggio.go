package main

import (
	"time"
  "math"
	"github.com/acruikshank/fmsynth"
)

const sampleRate = 44100
const decay = .05
const release = .25
const attack = .01
const duration = .5
const volume = .5

func main() {
	mixer := fmsynth.NewMixer(sampleRate)
	defer mixer.Terminate()

  go addNotes(mixer)


  mixer.Start()
	defer mixer.Stop()

	time.Sleep(10000 * time.Millisecond)
}

func addNotes(mixer *fmsynth.Mixer) {
  bass1 := &fmsynth.Algorithm1 {
    fmsynth.NewOscillator(   .5,  .99,    .2,   .2,   .25),
    fmsynth.NewOscillator(   .5,    1,   .01,    1,   .25),
    fmsynth.NewOscillator(    1,    2,   .01,    4,   .25),
    fmsynth.NewOscillator(   .5, 1.01,   .04,    1,   .25),
    fmsynth.NewOscillator(   .5,   .5,   .01,    4,   .25),
    fmsynth.NewOscillator(   .5,  .25,   .01,   16,   .25),
    0.2,
  }

  // bass1 := &fmsynth.Algorithm32 {
  //   fmsynth.NewOscillator(   1,    1,   .04,   .01,   .25),
  //   fmsynth.NewOscillator(  .75,   2,   .08,   .01,   .25),
  //   fmsynth.NewOscillator( .5,     4,   .08,   .01,   .25),
  //   fmsynth.NewOscillator(.125,    8,   .08,   .01,   .25),
  //   fmsynth.NewOscillator(.062,   16,   .10,   .02,   .25),
  //   fmsynth.NewOscillator(.016,   32,   .10,   .02,   .25),
  //   0,
  // }

  notes := []float64{
    freqForNote(33+12), // a
    freqForNote(31+12), // g
    freqForNote(29+12), // f
    freqForNote(27+12), // e flat
    freqForNote(26+12), // d
    freqForNote(24+12), // c
    freqForNote(22+12), // b flat
    // freqForNote(39), // e flat
    // freqForNote(43), // g
    // freqForNote(43), // g
    // freqForNote(29), // f
    // freqForNote(26), // d
    // freqForNote(29), // f
    // freqForNote(33), // a
    // freqForNote(29), // f
    // freqForNote(26), // d
    // freqForNote(24), // c
    // freqForNote(26), // d
    // freqForNote(39), // e flat
    // freqForNote(26), // d
    // freqForNote(39), // e flat
    // freqForNote(43), // g
    // freqForNote(33), // a
    // freqForNote(43), // g
    // freqForNote(39), // e flat
    // freqForNote(24), // c
    // freqForNote(26), // d
    // freqForNote(29), // f
    // freqForNote(33), // a
    // freqForNote(29), // f
    // freqForNote(26), // d
    // freqForNote(29), // f
    // freqForNote(26), // d
    // freqForNote(24), // c
  }
  clock := 0.0
  noteIndex := 0
  noteMod := 2 * len(notes) - 2

  for {
    noteCount := mixer.ActiveNotes()
    for i := 0; i < 6 - noteCount; i++ {
      var freq float64
      if noteIndex % noteMod >= len(notes) {
        freq = notes[2*len(notes) - (noteIndex % noteMod) - 2]
      } else {
        freq = notes[noteIndex % noteMod]
      }
      mixer.AddNote(fmsynth.NewNote(bass1, freq, clock, clock+duration, volume * math.Exp(50/math.Pow(freq,2) ) / math.E), false)
      clock += duration
      noteIndex++
    }
	  time.Sleep(time.Duration(float64(time.Second) * duration))
  }
}

func freqForNote(note int) float64 {
  return 440.0 * math.Pow(2.0, float64(note - 69) / 12.0)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
