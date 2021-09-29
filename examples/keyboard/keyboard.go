package main

import (
  "github.com/nsf/termbox-go"
  "fmt"
  "math"
  // "time"
	"github.com/acruikshank/fmsynth"
)

const sampleRate = 44100
const decay = .05
const release = .25
const attack = .01
const duration = 1
const volume = .5

func addNote(mixer *fmsynth.Mixer, note int) {
 bass1 := &fmsynth.Algorithm1 {
   fmsynth.NewOscillator(   .5,  .999,    .2,   .8,   .25),
   fmsynth.NewOscillator(   .5,    1,   .01,    1,   .25),
   fmsynth.NewOscillator(    .25,    2,   .05,    2,   .25),
   fmsynth.NewOscillator(   .5, 1.001,   .04,    4,   .25),
   fmsynth.NewOscillator(   .5,   1.5,   .01,    2,   .25),
   fmsynth.NewOscillator(   .5,  1.25,   .01,   4,   .25),
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

  freq := freqForNote( 36 + note )
  sound := fmsynth.NewNote(bass1, freq, 0, 0, volume * math.Exp(50/math.Pow(freq,4) ) / math.E)
  mixer.StopNoteWithDelay(sound, duration)
  mixer.AddNote(sound, true)
}

func freqForNote(note int) float64 {
  return 440.0 * math.Pow(2.0, float64(note - 69) / 12.0)
}

func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}

func drawBox(left int, top int, width int, height int, fg, bg termbox.Attribute) {
	for i := top; i < top+height; i++ {
    for j := left; j < left+width; j++ {
      termbox.SetCell(j, i, 0, fg, bg)
    }
  }
}

func drawBorder(left int, top int, width int, height int, fg, bg termbox.Attribute) {
  right := left + width - 1
  bottom := top + height - 1
  termbox.SetCell(left, top, 0x250C, fg, bg)
  termbox.SetCell(right, top, 0x2510, fg, bg)
  termbox.SetCell(left, bottom, 0x2514, fg, bg)
  termbox.SetCell(right, bottom, 0x2518, fg, bg)

  for i := left+1; i < right; i++ {
  	termbox.SetCell(i, top, 0x2500, fg, bg)
  	termbox.SetCell(i, bottom, 0x2500, fg, bg)
  }
  for i := top+1; i < bottom; i++ {
  	termbox.SetCell(left, i, 0x2502, fg, bg)
  	termbox.SetCell(right, i, 0x2502, fg, bg)
  }
}

func keyToNote(key rune) int {
  if key == 97 { return 0 }
  if key == 119 { return 1 }
  if key == 115 { return 2 }
  if key == 101 { return 3 }
  if key == 100 { return 4 }
  if key == 102 { return 5 }
  if key == 116 { return 6 }
  if key == 103 { return 7 }
  if key == 121 { return 8 }
  if key == 104 { return 9 }
  if key == 117 { return 10 }
  if key == 106 { return 11 }
  if key == 107 { return 12 }
  return -1
}

func draw_keyboard(note int) {
  printf_tb(56, 15, termbox.ColorWhite, termbox.ColorBlack, "Press Q to quit")
  backgrounds := []termbox.Attribute {
    termbox.ColorWhite,
    termbox.ColorBlack,
    termbox.ColorWhite,
    termbox.ColorBlack,
    termbox.ColorWhite,
    termbox.ColorWhite,
    termbox.ColorBlack,
    termbox.ColorWhite,
    termbox.ColorBlack,
    termbox.ColorWhite,
    termbox.ColorBlack,
    termbox.ColorWhite,
    termbox.ColorWhite,
  }

  if note >= 0 {
    backgrounds[note] = termbox.ColorYellow
  }

  termbox.HideCursor()
  drawBox(0, 0, 8, 14, termbox.ColorWhite, backgrounds[0])
  drawBox(9, 0, 8, 14, termbox.ColorWhite, backgrounds[2])
  drawBox(18, 0, 8, 14, termbox.ColorWhite, backgrounds[4])
  drawBox(27, 0, 8, 14, termbox.ColorWhite, backgrounds[5])
  drawBox(36, 0, 8, 14, termbox.ColorWhite, backgrounds[7])
  drawBox(45, 0, 8, 14, termbox.ColorWhite, backgrounds[9])
  drawBox(54, 0, 8, 14, termbox.ColorWhite, backgrounds[11])
  drawBox(63, 0, 8, 14, termbox.ColorWhite, backgrounds[12])

  drawBox(5, -1, 7, 8, termbox.ColorWhite, backgrounds[1])
  drawBorder(5, -1, 7, 8, termbox.ColorWhite, backgrounds[1])

  drawBox(14, -1, 7, 8, termbox.ColorWhite, backgrounds[3])
  drawBorder(14, -1, 7, 8, termbox.ColorWhite, backgrounds[3])

  drawBox(32, -1, 7, 8, termbox.ColorWhite, backgrounds[6])
  drawBorder(32, -1, 7, 8, termbox.ColorWhite, backgrounds[6])

  drawBox(41, -1, 7, 8, termbox.ColorWhite, backgrounds[8])
  drawBorder(41, -1, 7, 8, termbox.ColorWhite, backgrounds[8])

  drawBox(50, -1, 7, 8, termbox.ColorWhite, backgrounds[10])
  drawBorder(50, -1, 7, 8, termbox.ColorWhite, backgrounds[10])

	termbox.SetCell(3, 12, 'A', termbox.ColorBlack|termbox.AttrBold, backgrounds[0]) // 97
	termbox.SetCell(12, 12, 'S', termbox.ColorBlack|termbox.AttrBold, backgrounds[2]) // 115
	termbox.SetCell(21, 12, 'D', termbox.ColorBlack|termbox.AttrBold, backgrounds[4]) // 100
	termbox.SetCell(30, 12, 'F', termbox.ColorBlack|termbox.AttrBold, backgrounds[5]) // 102
	termbox.SetCell(39, 12, 'G', termbox.ColorBlack|termbox.AttrBold, backgrounds[7]) // 103
	termbox.SetCell(48, 12, 'H', termbox.ColorBlack|termbox.AttrBold, backgrounds[9]) // 104
	termbox.SetCell(57, 12, 'J', termbox.ColorBlack|termbox.AttrBold, backgrounds[11]) // 106
	termbox.SetCell(66, 12, 'K', termbox.ColorBlack|termbox.AttrBold, backgrounds[12]) // 107

	termbox.SetCell(8, 5, 'W', termbox.ColorWhite|termbox.AttrBold, backgrounds[1]) // 119
	termbox.SetCell(17, 5, 'E', termbox.ColorWhite|termbox.AttrBold, backgrounds[3]) // 101
	termbox.SetCell(35, 5, 'T', termbox.ColorWhite|termbox.AttrBold, backgrounds[6]) // 116
	termbox.SetCell(44, 5, 'Y', termbox.ColorWhite|termbox.AttrBold, backgrounds[8]) // 121
	termbox.SetCell(53, 5, 'U', termbox.ColorWhite|termbox.AttrBold, backgrounds[10]) // 117

}

func main() {
  mixer := fmsynth.NewMixer(sampleRate)
	defer mixer.Terminate()

  mixer.Start()
	defer mixer.Stop()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	draw_keyboard(-1)
	termbox.Flush()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC || ev.Ch == 113 {
				break loop
			}

      note := keyToNote(ev.Ch)
      if note >= 0 {
        addNote(mixer, note)
      }

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_keyboard(note)

			termbox.Flush()
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw_keyboard(-1)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
