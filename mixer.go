package fmsynth

import (
	"github.com/gordonklaus/portaudio"
  "sync"
  // "time"
  // "fmt"
)

type Mixer struct {
	*portaudio.Stream
  notes []*Note
  sampleRate float64
  clock float64
  *sync.RWMutex
}

func NewMixer(sampleRate float64) *Mixer {
  portaudio.Initialize()
	mixer := &Mixer{nil, make([]*Note,0), sampleRate, 0, &sync.RWMutex{}}
	var err error
	mixer.Stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, mixer.processAudio)
	if err != nil {
		panic(err)
  }
  return mixer
}

func (m *Mixer) Terminate() {
  portaudio.Terminate()
}

func (m *Mixer) ActiveNotes() int {
  m.Lock()
  defer m.Unlock()

  m.cleanNotes()
  return len(m.notes)
}

func (m *Mixer) AddNote(note *Note, start bool) {
  m.Lock()
  defer m.Unlock()

  m.cleanNotes()
  m.notes = append(m.notes, note)
  if start {
    note.start = m.clock
  }
}

func (m *Mixer) StartNote(note *Note) {
  m.Lock()
  defer m.Unlock()

  note.start = m.clock
}

func (m *Mixer) StopNoteWithDelay(note *Note, delay float64) {
  m.Lock()
  defer m.Unlock()

  note.end = m.clock + delay
}

func (m *Mixer) StopNote(note *Note) {
   m.StopNoteWithDelay(note, 0.0)
}

func (m *Mixer) cleanNotes() {
  next := make([]*Note, 0, 8)
  for _, note := range m.notes {
    if (! note.complete(m.clock)) {
      next = append(next, note)
    }
  }
  m.notes = next
}

func (m *Mixer) processAudio(out [][]float32) {
  m.RLock()
  defer m.RUnlock()

  inc := 1.0/m.sampleRate

	for i := range out[0] {
    amplitude := 0.0
    for _, note := range m.notes {
      amplitude += note.play(m.clock)
    }
  	out[0][i] = float32(amplitude)
    m.clock += inc
	}
}
