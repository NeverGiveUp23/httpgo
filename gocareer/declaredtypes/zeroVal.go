package declaredtypes

// this initial value is 0, just like in most languages if it isn't initialized
var zeroInit int

// go literals -> use underscores for representation of a thousand
var thousand = 1_234

var flag bool // no value assigned, set to false (0)
var isAwesome = true

// uint means unsigned or non-negative integers
var unsignedInt8 uint8 = 255                   // uint8 0 - 255 (1 byte)
var unsignedInt16 uint16 = 65535               // unint16 -> 0 -65535
var unsignedInt32 uint32 = 4294967295          // unint32 -> 0 - 4294967295
var unsignedInt64 uint64 = 1844674407370955165 // unint64 -> 0 -18446744073709551615

var x = 9
var y = 9

var addTogether = x + y

var myFirstInitial rune = 'F'

var i, l int
var num1, num2 = 10, 30

// declare multiple variables like so:
var (
	f, g string
	z    = 30
	k    int
)

func operator() string {

	// := -> has one limitation, it can only work in the function declaration
	// if you're using it outside the function you need to use var
	d := f
	t, u := 30, "Hello"
	print(d, t)
	return u
}
