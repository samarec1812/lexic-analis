package analysator

import (
	"unicode"
)

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func SplitText(text string) ([]string, []string) {
	/* words := regexp.MustCompile("[\t\n]{1}").Split(text, -1)
	for index, word := range words {
		fmt.Println(strconv.Itoa(index) + ": " + word)
	}
	return words */
	res := make([]string, 0)
	word := ""
	word2 := ""
	flag := false
	for index, s := range text {
		if (s == '>' || s == '<') && text[index+1] == '='  {
			flag = true
			word2 += string(s)
		} else if s == '=' && flag {
			word2 += string(s)
			res = append(res, word2)
			flag = false
			word2 = ""

		} else if unicode.IsLetter(s) || unicode.IsDigit(s) {
			word += string(s)
		} else {
			res = append(res, word)
			if !(s == ' ' ||  s == '\t' || s == '\r') {
				if (s == '>' && text[index+1] == '=') || (s == '<' && text[index+1] == '=') { continue }
				res = append(res, string(s))
			}
			word = ""
		}
	}
	filtered := Filter(res, func(w string) bool {
		if w == "\t" || w == "\n" || w == "" {
			return false
		}
		return true
	})

	filtered2 := Filter(res, func(w string) bool {
		if w == "\t" || w == "" {
			return false
		}
		return true
	})
	return filtered, filtered2
}
var cur int = 0
func Checker(text []string) bool {
	if cur >= len(text) {
		return false
	}
	if text[cur] == "program" {
		cur += 1
		return ProgramName(text)
	} else {
		return false
	}
}

// имя программы
func ProgramName(text []string) bool {
	if cur >= len(text) { return false }
	if Identificator(text[cur]) {
		cur += 1
		if cur >= len(text) { return false }
		if text[cur] == ";" {
			cur += 1
			return Block(text)
		}
	}
	return false
}

// проверка идентификатора
func Identificator(text string) bool {
	if !unicode.IsLetter(rune(text[0])) { return false }
	for _, value := range text {
		if !(unicode.IsLetter(value) || unicode.IsDigit(value)) { return false }
	}
	return true
}

// Блок программы
func Block(text []string) bool {
	return DeclarativePart(text)
}

// Блок описания
func DeclarativePart(text []string) bool {
	return VarDeclaration(text)
}

// Блок описания переменных
func VarDeclaration(text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == "var" {
		cur += 1
		return VarList(text)
	} else {
		return FunDeclaration(text)
	}
}
// Блок описания переменных
func VarList(text []string) bool {
	if IDList(text) {
		return FunDeclaration(text)
	} else {
		return false
	}
}

// Список глобальных переменных
func IDList(text []string) bool {
	if cur >= len(text) { return false }
	searchID := true
	searchDoublePoint := false
	searchComma := false
	searchType := false
	searchPointwithComma := false
	for cur < len(text) && (text[cur] != "begin" && text[cur] != "function") {
		if searchID {
			if Identificator(text[cur]) {
				searchID = false
				searchComma = true
				searchDoublePoint = true
			} else { return false }
		} else if searchComma || searchDoublePoint {
			if text[cur] == "," {
				searchComma = false
				searchDoublePoint = false
				searchID = true
			} else if text[cur] == ":" {
				searchComma = false
				searchDoublePoint = false
				searchType = true
			} else { return false }
		} else if searchType {
			if !TypeID(text[cur]) {
				return false
			}
			searchType = false
			searchPointwithComma = true
		} else if searchPointwithComma {
			if !(text[cur] == ";"){
				return false
			}
			searchID = true
			searchPointwithComma = false
		} else {
			return false
		}
		cur += 1
	}
	if searchID && cur < len(text) && (text[cur] == "begin" || text[cur] == "function") {
		return true
	} else {
		return false
	}
}


// Тип переменной
func TypeID(text string) bool {
	return text == "integer" || text == "boolean"
}

// Блок описания функций
func FunDeclaration(text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == "function" {
		cur += 1
		return FunName(text)
	} else {
		return Body(text)
	}
}

// Имя функции
func FunName(text []string) bool {
	if cur >= len(text) { return false }
	if Identificator(text[cur]) {
		cur += 1
		return FunParam(text)
	} else {
		return false
	}
}

// Блок описания параметров функции
func FunParam(text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == "(" {
		cur += 1
		if FunIDList(text) {
			cur += 1
			return FunType(text)
		} else {
			return false
		}
	} else {
		return FunType(text)
	}
}

// Параметры функции
func FunIDList(text []string) bool {
	if cur >= len(text) { return false }
	searchID := true
	searchDoublePoint := false
	searchComma := false
	searchType := false
	searchPointWithComma := false
	for cur < len(text) && text[cur] != ")" {
		if searchID {
			if Identificator(text[cur]) {
				searchID = false
				searchComma = true
				searchDoublePoint = true
			} else { return false }
		} else if searchComma || searchDoublePoint {
			if text[cur] == "," {
				searchComma = false
				searchID = true
				searchDoublePoint = false
			} else if text[cur] == ":" {
				searchComma = false
				searchDoublePoint = false
				searchType = true
			} else {
				return false
			}
		} else if searchType {
			if !TypeID(text[cur]) { return false }
			searchType = false
			searchPointWithComma = true
		} else if searchPointWithComma {
			if !(text[cur] == ";") {
				return false
			}
			searchID = true
			searchPointWithComma = false
		} else {
			return false
		}
		cur += 1
	}
	if (cur < len(text) && (text[cur] == ")") && (TypeID(text[cur-1]) || text[cur-1] == "(")) {
		return true
	} else {
		return false
	}
}

// Тип функции
func FunType(text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == ":" {
		cur += 1
		if cur >= len(text) {
			return false
		} else {
			if (TypeID(text[cur]) && cur + 1 < len(text) && text[cur+1] == ";") {
				cur += 2
				return FunDeclarativePart(text)
			} else {
				return false
			}
		}
	} else { return false }
}

// Описание блока переменных var в функции
func FunDeclarativePart(text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == "var" {
		cur += 1
		if !FunIdList(text) { return false }
	}
	if text[cur] == "begin" {
		cur += 1
		if FunBlock(text) {
			return FunDeclaration(text)
		}
	}
	return false
}

// Список переменных в блоке функции описания переменных var
func FunIdList(text []string) bool {
	if cur >= len(text) { return false }
	searchID := true
	searchDoublePoint := false
	searchComma := false
	searchType := false
	searchPointWithComma := false
	for cur < len(text) && text[cur] != "begin" {
		if searchID {
			if Identificator(text[cur]) {
				searchID = false
				searchComma = true
				searchDoublePoint = true
			} else {
				return false
			}
		} else if searchComma || searchDoublePoint {
			if text[cur] == "," {
				searchComma = false
				searchID = true
				searchDoublePoint = false
			} else if text[cur] == ":" {
				searchComma = false
				searchType = true
				searchDoublePoint = false
			} else {
				return false
			}
		} else if searchType {
			if !TypeID(text[cur]) {
				return false
			}
			searchType = false
			searchPointWithComma = true
		} else if searchPointWithComma {
			if !(text[cur] == ";") {
				return false
			}
			searchID = true
			searchPointWithComma = false
		} else {
			return false
		}
		cur += 1
	}
	if searchID && cur < len(text) && text[cur] == "begin" {
		return true
	} else {
		return false
	}
}

// Обработка тела функции
func FunBlock(text []string) bool {
	if cur >= len(text) { return false }
	if Stmt(text) {
		if text[cur] == "end" {
			cur += 1
			if cur >= len(text) { return false }
			if text[cur] == ";" {
				cur += 1
				return true
			}
		}
	}
	return false
}

// Обработка главного тела begin ... end
func Body (text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == "begin" {
		cur += 1
		if Stmt(text) {
			if text[cur] == "end" {
				cur += 1
				if cur >= len(text) { return false }
				if text[cur] == "." && (cur + 1 == len(text)) {
					return true
				}
			}
		}
	}
	return false
}

// Список утвержений
func Stmt(text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == "read" || text[cur] == "write" {
		cur += 1
		if ReadAndWriteStmt(text) {
			return Stmt(text)
		}
	} else if text[cur] == "for" {
		cur += 1
		if ForStmt(text) {
			return Stmt(text)
		}
	} else if text[cur] == "if" {
		cur += 1
		if IfStmt(text) {
			return Stmt(text)
		}
	} else if text[cur] == "end" {
		return true
	} else if Identificator(text[cur]) {
		cur += 1
		if cur >= len(text) { return false }
		if !(text[cur] == ":") { return false }
		cur += 1
		if cur >= len(text) { return false }
		if !(text[cur] == "=") { return false }
		cur += 1
		curAfterID := cur
		if !Expr(text) {
			cur = curAfterID
			if CallFunc(text) {
				return Stmt(text)
			}
		} else {
			return Stmt(text)
		}
	} else if text[cur] == ";" {
		cur += 1
		return Stmt(text)
	}
	return false
}

func ReadAndWriteStmt(text []string) bool {
	if cur >= len(text) { return false }
	if text[cur] == "(" {
		searchID := true
		searchComma := false
		cur += 1
		for cur < len(text) && text[cur] != ")" {
			if searchID {
				if Identificator(text[cur]){
					searchID = false
					searchComma = true
				} else { return false }
			} else if searchComma {
				if text[cur] == "," {
					searchComma = false
					searchID = true
				} else { return false }
			} else {
				return false
			}
			cur += 1
		}
		if cur < len(text) && searchComma && text[cur] == ")" {
			cur += 1
			if cur >= len(text) { return false }
			if text[cur] == ";" {
				cur += 1
				return true
			}
		}
	}
	return false
}

// Проверка цикла for
func ForStmt(text []string) bool {
	if cur >= len(text) { return false }
	if !Identificator(text[cur])  { return false }
	cur += 1
	if cur >= len(text) { return false }
	if !(text[cur] == ":") { return false }
	cur += 1
	if cur >= len(text) { return false }
	if !(text[cur] == "=") { return false }
	cur += 1
	if cur >= len(text) { return  false }
	expr := make([]string, 0, 0)
	for cur < len(text) && text[cur] != "to" && text[cur] != "downto" {
		expr = append(expr, text[cur])
		cur += 1
	}
	if !CheckArifmExpr(expr) { return false }
	if cur >= len(text) || !(text[cur] == "to" || text[cur] == "downto") { return false }
	expr = []string{}
	for cur < len(text) && text[cur] != "do" {
		expr = append(expr, text[cur])
		cur += 1
	}
	if !CheckArifmExpr(expr) { return false }
	cur += 1
	if cur >= len(text) { return false }
	if text[cur] == "begin" {
		cur += 1
		return FunBlock(text)
	} else {
		return Stmt(text)
	}
}

// Проверка условного оператора
func IfStmt(text[]string) bool {
	if cur >= len(text) { return false }
	expr := make([]string, 0, 0)
	for cur < len(text) && text[cur] != "then" {
		expr = append(expr, text[cur])
		cur += 1
	}
	if !CheckLogExpr(expr) { return false }
	if cur >= len(text) || text[cur] != "then" { return false }
	cur += 1
	if cur >= len(text) { return false }
	if text[cur] == "begin" {
		cur += 1
		return FunBlock(text)
	} else {
		return Stmt(text)
	}
}

// Выделение арифметического выражения
func Expr(text []string) bool {
	if cur >= len(text) { return false }
	expr := make([]string, 0, 0 )
	for cur < len(text) && text[cur] != ";" {
		expr = append(expr, text[cur])
		cur += 1
	}
	if CheckArifmExpr(expr) {
		if cur < len(text) && text[cur] == ";" {
			cur += 1
			return true
		}
	}
	return false
}

// Проверка логического выражения
func CheckLogExpr(expr []string) bool {
	balance := 0
	if len(expr) == 0 { return false }
	if len(expr) == 1 && (IsArifmOperation(expr[0]) || IsCondOperation(expr[0])) { return false }
	if expr[0] == "(" { balance += 1 }
	if expr[0] == ")" { return false }
	for i := 1; i < len(expr) - 1; i++ {
		if expr[i] == "(" {
			balance += 1
			if !(IsCondOperation(expr[i-1]) || expr[i - 1] == "(") { return false }
			if !(Identificator(expr[i+1]) || Number(expr[i+1]) || expr[i+1] == "(") { return false }
		} else if expr[i] == ")" {
			if balance > 0 {
				balance--
			} else { return false }
			if !(Identificator(expr[i-1]) || expr[i-1] == ")" || Number(expr[i-1])) {return false }
			if !(IsCondOperation(expr[i+1]) || expr[i+1] == ")") { return false }
		} else if IsLogOperation(expr[i]) {
			if !(Identificator(expr[i-1]) || Number(expr[i-1])) { return false }
			if !(Identificator(expr[i+1]) || Number(expr[i+1])) { return false }
		} else if IsCondOperation(expr[i]) {
			if expr[i-1] != ")" || expr[i+1] != "(" { return false }
		} else if Identificator(expr[i]) {
			if Identificator(expr[i-1]) || Identificator(expr[i+1]) { return false }
			if Number(expr[i-1]) || Number(expr[i+1]) { return false }
		} else if Number(expr[i]) {
			if Identificator(expr[i-1]) || Identificator(expr[i+1]) {return false}
			if Number(expr[i-1]) || Number(expr[i+1]) { return false }
		}
	}
	if !(Identificator(expr[len(expr) - 1]) || Number(expr[len(expr) - 1]) || expr[len(expr) - 1] == ")") {
		return false
	}
	if expr[len(expr) - 1] == ")" {balance--}
	if balance > 0 { return false }
	return true
}

// Проверка на операции сравнения
func IsLogOperation(ch string) bool {
	return ch == "=" || ch == "<" || ch == ">" || ch == ">=" || ch == "<="
}

// Проверка на логические операнды
func IsCondOperation(ch string) bool {
	return ch == "and" || ch == "or"
}

// Проверка арифметического выражения
func CheckArifmExpr(expr []string) bool {
	balance := 0
	if len(expr) == 0 { return false }
	if IsArifmOperation(expr[0]) && expr[0] != "-" { return false }
	if len(expr) == 1 && IsArifmOperation(expr[0]) { return false }
	if expr[0] == "(" { balance += 1}
	if expr[0] == ")" { return false }

	for i := 1; i < len(expr) -1; i++ {
		if expr[i] == "(" {
			balance += 1
			if !IsArifmOperation(expr[i-1]) && !(expr[i-1] == "(") {

				return false }
			if !(Identificator(expr[i+1]) || expr[i+1] == "-" || Number(expr[i+1]) || expr[i+1] == "(") { return false }
		} else if expr[i] == ")" {
			if balance > 0 {
				balance--
			} else {
				return false
			}
			if !(Identificator(expr[i-1]) || expr[i-1] == ")" || Number(expr[i-1])) {
				return false
			}
			if !IsArifmOperation(expr[i+1]) && !(expr[i+1] == ")") {
				return false
			}
		} else if expr[i] == "+" {
				if !(Identificator(expr[i-1])  || Number(expr[i-1])) { return false }
				if !(Identificator(expr[i+1])  || Number(expr[i+1])) { return false }
			} else if expr[i] == "-" {
				if !(Identificator(expr[i-1]) || Number(expr[i-1])) { return false }
				if !(Identificator(expr[i+1]) || Number(expr[i+1])) { return false }
			} else if expr[i] == "*" {
				if !(Identificator(expr[i-1]) || Number(expr[i-1])) { return false }
				if !(Identificator(expr[i+1]) || Number(expr[i+1])) { return false }
			} else if expr[i] == "/" {
				if !(Identificator(expr[i-1]) || Number(expr[i-1])) { return false }
				if !(Identificator(expr[i+1]) || Number(expr[i+1])) { return false }


			//} else if expr[i] == "+" {
		//	if !(Identificator(expr[i-1]) || expr[i-1] == ")" || Number(expr[i-1])) { return false }
		//	if !(Identificator(expr[i+1]) || expr[i+1] ==" (" || Number(expr[i+1])) { return false }
		//} else if expr[i] == "-" {
		//	if !(Identificator(expr[i-1]) || expr[i-1] == "(" || expr[i-1] == ")" || Number(expr[i-1])) { return false }
		//	if !(Identificator(expr[i+1]) || expr[i+1] ==" (" || Number(expr[i+1])) { return false }
		//} else if expr[i] == "*" {
		//	if !(Identificator(expr[i-1]) || expr[i-1] == ")" || Number(expr[i-1])) { return false }
		//	if !(Identificator(expr[i+1]) || expr[i+1] ==" (" || Number(expr[i+1])) { return false }
		//} else if expr[i] == "/" {
		//	if !(Identificator(expr[i-1]) || expr[i-1] == ")" || Number(expr[i-1])) { return false }
		//	if !(Identificator(expr[i+1]) || expr[i+1] ==" (" || Number(expr[i+1])) { return false }
		} else if Identificator(expr[i]) {
			if Identificator(expr[i-1]) || Identificator(expr[i+1]) { return false }
			if Number(expr[i-1]) || Number(expr[i+1]) { return false }
		} else if Number(expr[i]) {
			if Identificator(expr[i-1]) || Identificator(expr[i+1]) { return false }
			if Number(expr[i-1]) || Number(expr[i+1]) { return false }
		}
	}
	if !(Identificator(expr[len(expr) - 1]) || Number(expr[len(expr) - 1]) || expr[len(expr) - 1] == ")") {
		return false
	}
	if expr[len(expr) - 1] == ")" { balance--}
	if balance > 0 { return false }
	return true
}

// Проверка символа на арифметическую операцию
func IsArifmOperation(ch string) bool {
	return ch == "+" || ch == "-" || ch == "*" || ch == "/"
}

// Проверка на вызов функции
func CallFunc(text []string) bool {
	if cur >= len(text) { return false }
	if !Identificator(text[cur]) { return false }
	cur += 1
	if text[cur] == "(" {
		cur += 1
		searchID := true
		searchComma := false
		for cur < len(text) && text[cur] != ")" {
			if searchID {
				if Identificator(text[cur]) || Number(text[cur]) {
					searchID = false
					searchComma = true
				} else {
					return false
				}
			} else if searchComma {
				if text[cur] == "," {
					searchComma = false
					searchID = true
				} else {
					return false
				}
			}
			cur += 1
		}
		if text[cur] == ")" && (searchComma || text[cur - 1] == "(") {
			cur += 1
			if cur >= len(text) { return false }
		}
	}
	if text[cur] == ";" {
		cur += 1
		return true
	}
	return false
}


// Проверка на число
func Number(text string) bool {
	for _, value := range text {
		if !unicode.IsDigit(value) {
			return false
		}
	}
	return true
}

func GetErrorLine(text []string) int {
	index := 0
	index2 := 0
	line := 1
	for index < cur {
		if text[index2] == "\n" {
			line += 1
		} else {
			index += 1
		}
		index2 += 1
	}
	return line
}
