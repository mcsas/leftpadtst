package cycle

//import "github.com/jonbodner/go3hours/s4/my-packages/leftpad-cycle"
import "github.com/medero/leftpad-cycle"

var DEFAULT_CHAR = ' '

func FormatDouble(s string, i int) string {
	return leftpad.Format(s+s, i)
}
