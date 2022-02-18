package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math";

const (
	SidesTriangle SidesNum = 3;
	SidesSquare SidesNum = 4;
	SidesCircle SidesNum = 0;
)

type SidesNum int;

func CalcSquare(sideLen float64, sidesNum SidesNum) float64 {
	var square float64;

	switch sidesNum {
    case 0:
        square = math.Pi*math.Pow(sideLen, 2);
    case 3:
        square = math.Sqrt(3) * math.Pow(sideLen, 2) / 4;
    case 4:
       square = math.Pow(sideLen, 2);
	}

	return square;
}
