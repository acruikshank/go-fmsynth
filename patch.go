package fmsynth

import (
  "math"
  // "time"
  // "fmt"
)

type Oscillator struct {
  envelope Envelope
  volume float64
  frequencyRatio float64
}

func NewOscillator(volume float64, frequencyRatio float64, attack float64, decay float64, release float64) Oscillator {
  return Oscillator { NewEnvelope(attack, decay, release), volume, frequencyRatio }
}

func (o *Oscillator) Value(baseFrequency float64, clock float64, start float64, end float64, modulation float64) float64 {
  return o.volume * o.envelope.amplitude(clock, start, end) * math.Sin(2*math.Pi*(baseFrequency * o.frequencyRatio * clock + modulation));
}

func (a *Oscillator) Complete(clock float64, end float64) bool {
  return clock < end + a.envelope.release
}

type Envelope struct {
  attack float64
  decay float64
  release float64
}

func NewEnvelope(attack float64, decay float64, release float64) Envelope {
  return Envelope {
    attack,
    decay,
    release,
  }
}

func (e *Envelope) amplitude(clock float64, start float64, end float64) float64 {
  if (clock < start) {
    return 0
  }

  noteClock := clock - start
  if noteClock < e.attack {  // attack
    return noteClock / e.attack;
  }

  if clock < end {  // decay
    return e.dampened(noteClock)
  }

  afterTime := clock - end
  if afterTime < e.release {  // release
    return mapValue(0, e.release, math.Min(noteClock / e.attack, e.dampened(end - start)), 0, afterTime)
  }

  return 0  // after note
}

func (e *Envelope) dampened(noteClock float64) float64 {
  dampen := math.Pow(e.decay, 2)
  return math.Pow(math.Max(1 - (noteClock - e.attack),0), dampen)
}

// func mapValue(a float64, b float64, c float64, d float64, x float64) float64 {
//   return c + (x - a) * (d - c) / (b - a)
// }
