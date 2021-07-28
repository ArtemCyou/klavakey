//преобразует данные введенные пользователем с клавиатуры в цифры
package klavkey

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// преобразует введенные данные в дробное число
func GetFloat() (float64, error) {
	reader:= bufio.NewReader(os.Stdin)
	input, err:= reader.ReadString('\n')
	if err!=nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err:= strconv.ParseFloat(input, 64)
	if err!=nil{
		return 0, err
	}
	return number, nil
	}

// преобразует введенные данные в целое число
func GetInt() (int, error)  {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return number, nil
}