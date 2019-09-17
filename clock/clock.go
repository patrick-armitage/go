package clock

import(
  "strconv"
)

// Clock type stores hour and minute of a dateless clock
type Clock struct {
  hour, minute int
}

// New returns a clock object with properly calculated hours and minutes
// based on the provided input in {00-23}:{00-59} format
func New(hr int, min int) (newClock Clock) {
  for min > 59 {
    min = min - 60
    hr = hr + 1
  }

  for min < 0 {
    min = min + 60
    hr = hr - 1
  }

  for hr > 23 {
    hr = hr - 24
  }

  for hr < 0 {
    hr = hr + 24
  }

  newClock = Clock{
    hour: hr,
    minute: min,
  }

  return newClock
}

// Add modifies Clock type to add integer minutes to current clock object's time
func (c Clock) Add(min int) Clock {
  c.minute = c.minute + min

  for c.minute > 59 {
    c.minute = c.minute - 60
    c.hour = c.hour + 1

    if c.hour > 23 {
      c.hour = c.hour - 24
    }
  }

  return c
}

// Subtract modifies Clock type to subtract integer minutes from clock object's time
func (c Clock) Subtract(min int) Clock {
  c.minute = c.minute - min

  for c.minute < 0 {
    c.minute = c.minute + 60
    c.hour = c.hour - 1

    if c.hour < 0 {
      c.hour = c.hour + 24
    }
  }

  return c
}

// String modifies Clock type to return string format of clock object's time
// in the format "{00-23}:{00-59}"
func (c Clock) String() string {
  hourStr := ""
  if c.hour < 10 {
    hourStr = "0"
  }
  hourStr = hourStr + strconv.Itoa(c.hour)
  minStr := ""
  if c.minute < 10 {
    minStr = "0"
  }
  minStr = minStr + strconv.Itoa(c.minute)

  clockStr := hourStr + ":" + minStr

  return clockStr
}
