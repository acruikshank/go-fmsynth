package fmsynth

import (
  "math"
  // "time"
  // "fmt"
)

type Algorithm interface {
  Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64)
  Complete(clock float64, end float64) bool
}

type Algorithm1 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm1) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, o5)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, o4)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, 0)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  // fmt.Printf("%f %f\n", a.Feedback * feedback, o6)
  return o1 + o3, o6
}
func (a *Algorithm1) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc4.envelope.release, a.Osc1.envelope.release)
}

type Algorithm2 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm2) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, 0)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, o5)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, o4)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o3, o2
}
func (a *Algorithm2) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc4.envelope.release, a.Osc1.envelope.release)
}

type Algorithm3 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm3) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, o5)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, 0)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, o3)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o4, o6
}
func (a *Algorithm3) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc4.envelope.release, a.Osc1.envelope.release)
}

type Algorithm4 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm4) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, o5)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, 0)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, o3)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o4, o4
}
func (a *Algorithm4) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc4.envelope.release, a.Osc1.envelope.release)
}

type Algorithm5 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm5) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, 0)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, o4)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, 0)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o3 + o5, o6
}
func (a *Algorithm5) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc5.envelope.release, math.Max(a.Osc3.envelope.release, a.Osc1.envelope.release))
}

type Algorithm6 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm6) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, 0)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, o4)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, 0)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o3 + o5, o5
}
func (a *Algorithm6) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc5.envelope.release, math.Max(a.Osc3.envelope.release, a.Osc1.envelope.release))
}

type Algorithm7 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm7) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, 0)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, o4 + o5)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, 0)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o3, o6
}
func (a *Algorithm7) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc3.envelope.release, a.Osc1.envelope.release)
}

type Algorithm8 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm8) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, 0)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, o4 + o5)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, 0)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o3, o6
}
func (a *Algorithm8) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc3.envelope.release, a.Osc1.envelope.release)
}

type Algorithm9 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm9) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, 0)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, o6)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, 0)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, o4 + o5)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, o2)
  return o1 + o3, o2
}
func (a *Algorithm9) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc3.envelope.release, a.Osc1.envelope.release)
}


type Algorithm32 struct { Osc1 Oscillator; Osc2 Oscillator; Osc3 Oscillator; Osc4 Oscillator; Osc5 Oscillator; Osc6 Oscillator; Feedback float64 }
func (a *Algorithm32) Value(baseFrequency float64, clock float64, start float64, end float64, feedback float64) (float64, float64) {
  o6 := a.Osc6.Value(baseFrequency, clock, start, end, a.Feedback * feedback)
  o5 := a.Osc5.Value(baseFrequency, clock, start, end, 0)
  o4 := a.Osc4.Value(baseFrequency, clock, start, end, 0)
  o3 := a.Osc3.Value(baseFrequency, clock, start, end, 0)
  o2 := a.Osc2.Value(baseFrequency, clock, start, end, 0)
  o1 := a.Osc1.Value(baseFrequency, clock, start, end, 0)
  return o1 + o2 + o3 + o4 + o5 + o6, o2
}
func (a *Algorithm32) Complete(clock float64, end float64) bool {
  return clock > end + math.Max(a.Osc1.envelope.release,
    math.Max(a.Osc2.envelope.release,
      math.Max(a.Osc3.envelope.release,
        math.Max(a.Osc4.envelope.release,
          math.Max(a.Osc5.envelope.release, a.Osc6.envelope.release)))));
}
