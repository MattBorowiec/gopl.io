package units

import (
	"fmt"
)

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%.4g feet", f) }
func (f Meters) String() string { return fmt.Sprintf("%.4g meters", f) }

type Pounds float64
type Kilograms float64

func (p Pounds) String() string    { return fmt.Sprintf("%.4g lb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%.4g kg", k) }

func PToK(p Pounds) Kilograms { return Kilograms(p * 0.45359237) }
func KToP(k Kilograms) Pounds { return Pounds(k / 0.45359237) }

func FToM(f Feet) Meters { return Meters(f / 3.2808) }
func MToF(m Meters) Feet { return Feet(m * 3.2808) }
