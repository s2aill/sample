package main
import(
	"fmt"
	"strconv"
	
	"github.com/pkg/errors"

)
func main () {
	toInt , err := toInt ("Gopher")
	if err != nil {
		fmt . Println (err)
		return
	}
	fmt.Println(toInt)
}

func toInt (value string) (int , error) {
	i, err := strconv . Atoi ( value )
	if err != nil {
		return 0, errors.Wrap(err , "•Ï Š· ‚É Ž¸ ”s")
	}
	return i, nil
	
}