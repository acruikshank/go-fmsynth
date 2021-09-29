package fmsynth

type Note struct {
  patch Algorithm
  frequency float64
  start float64
  end float64
  volume float64
  feedbackValue float64
}

func NewNote(patch Algorithm, frequency float64, start float64, end float64, volume float64) *Note {
  return &Note { patch, frequency, start, end, volume, 0 }
}

func (note *Note) play(clock float64) float64 {
  value, feedbackValue := note.patch.Value(note.frequency, clock, note.start, note.end, note.feedbackValue)
  note.feedbackValue = feedbackValue
  return note.volume * value
}

func (note *Note) complete(clock float64) bool {
  return note.patch.Complete(clock, note.end)
}

func mapValue(a float64, b float64, c float64, d float64, x float64) float64 {
  return c + (x - a) * (d - c) / (b - a)
}
