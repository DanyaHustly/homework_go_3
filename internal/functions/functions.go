package functions

import "errors"

type User struct {
	ID     int
	Name   string
	Active bool
}

// 01. SafeDivide выполняет целочисленное деление.
// TODO: при b == 0 верните 0 и ошибку "division by zero".
func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 02. FindUserByID ищет пользователя по ID.
// TODO: найденный пользователь возвращается с true; отсутствие — User{} и false.
func FindUserByID(users []User, id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

// 03. FindActiveUser ищет пользователя с нужным ID только среди активных.
// TODO: неактивный пользователь считается ненайденным.
func FindActiveUser(users []User, id int) (User, bool) {
	for _, user := range users {
		if user.Active && user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

// 04. SumAll складывает произвольное количество чисел.
// TODO: вызов без аргументов должен вернуть 0.
func SumAll(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// 05. Apply вызывает переданную функцию op для a и b.
// TODO: верните результат op без изменения аргументов.
func Apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 06. ApplyIf применяет transform к value только при allowed=true.
// TODO: при false верните исходное value и не вызывайте transform.
func ApplyIf(value int, allowed bool, transform func(int) int) int {
	if allowed {
		return transform(value)
	}
	return value
}

// 07. NewCounter создаёт замыкание-счётчик.
// TODO: первый вызов возвращённой функции должен вернуть start+1,
// каждый следующий — ещё на единицу больше.
func NewCounter(start int) func() int {
	return func() int {
		start++
		return start
	}
}

// 08. NewAccumulator создаёт замыкание с накопленной суммой.
// TODO: каждый вызов добавляет аргумент к текущему состоянию и возвращает сумму.
func NewAccumulator(initial int) func(int) int {
	return func(additional int) int {
		initial += additional
		return initial
	}
}

// 09. MakeMultiplier возвращает функцию умножения на factor.
// TODO: factor должен сохраняться внутри возвращённой функции.
func MakeMultiplier(factor int) func(int) int {
	return func(additional int) int {
		return factor * additional
	}
}

// 10. DeferOrder возвращает строку, показывающую порядок defer.
// TODO: тело добавляет "body", затем зарегистрированы defer "first" и "second".
// Итоговая строка должна быть "body-second-first".
func DeferOrder() (result string) {
	result = "body"
	defer func() {
		result = result + "-first"
	}()
	defer func() {
		result = result + "-second"
	}()
	return result
}

// 11. CaptureDeferArgument демонстрирует вычисление аргумента defer сразу.
// TODO: сначала value="first", затем зарегистрируйте defer с value как параметром,
// после регистрации поменяйте value на "second". Верните захваченное значение.
func CaptureDeferArgument() (result string) {
	value := "first"
	defer func(str string) {
		result = str
	}(value)
	value = "second"
	return result
}

// 12. ReadDeferredVariable демонстрирует чтение переменной deferred-замыканием.
// TODO: deferred-функция без параметров должна прочитать value после того,
// как оно изменилось с "first" на "second".
func ReadDeferredVariable() (result string) {
	value := "first"
	defer func() {
		result = value
	}()
	value = "second"
	return result
}

// 13. IncrementNamedResult возвращает value, увеличенный deferred-функцией на 1.
// TODO: используйте именованное возвращаемое значение.
func IncrementNamedResult(value int) (result int) {
	defer func() {
		result++
	}()
	return value
}

// 14. RunWithCleanup выполняет action, а cleanup откладывает до выхода.
// TODO: верните две строки в порядке фактического выполнения: action, cleanup.
func RunWithCleanup(action, cleanup func() string) (events []string) {
	defer func() {
		events = append(events, cleanup())
	}()
	events = append(events, action())
	return events
}

// 15. ChooseOperation возвращает функцию для "add", "sub" или "mul".
// TODO: для неизвестного имени верните nil и false.
func ChooseOperation(name string) (func(int, int) int, bool) {
	add := func(a, b int) int {
		return a + b
	}
	sub := func(a, b int) int {
		return a - b
	}
	mul := func(a, b int) int {
		return a * b
	}

	switch name {
	case "add":
		return add, true
	case "sub":
		return sub, true
	case "mul":
		return mul, true
	}
	return nil, false
}
