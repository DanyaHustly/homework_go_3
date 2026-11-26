package slices

// 01. First возвращает первый элемент и true.
// TODO: для nil и пустого слайса верните zero value и false.
func First(items []int) (int, bool) {
	if items == nil || len(items) == 0 {
		return 0, false
	}
	return items[0], true
}

// 02. Last возвращает последнюю строку и true.
// TODO: для nil и пустого слайса верните пустую строку и false.
func Last(items []string) (string, bool) {
	if items == nil || len(items) == 0 {
		return "", false
	}
	return items[len(items)-1], true
}

// 03. SliceInfo создаёт обычный срез items[low:high].
// TODO: верните сам срез, его len и cap. В тестах границы корректны.
func SliceInfo(items []int, low, high int) (part []int, length, capacity int) {

	return items[low:high], len(items[low:high]), cap(items[low:high])
}

// 04. FullSliceInfo создаёт полный срез items[low:high:max].
// TODO: верните сам срез, его len и cap. В тестах границы корректны.
func FullSliceInfo(items []int, low, high, max int) (part []int, length, capacity int) {
	return items[low:high:max], len(items[low:high:max]), cap(items[low:high:max])
}

// 05. ChangeFirst меняет первый элемент переданного слайса.
// TODO: пустой слайс оставьте без изменений.
func ChangeFirst(part []int, value int) {
	if len(part) == 0 {
		return
	}
	part[0] = value
}

// 06. MutateWindowInFunction создаёт срез items[low:high], передаёт его
// в ChangeFirst и возвращает исходный items.
// TODO: изменение должно быть видно в исходном слайсе. Пустое окно не меняйте.
func MutateWindowInFunction(items []int, low, high, value int) []int {
	ChangeFirst(items[low:high], value)
	return items
}

// 07. AppendWindowInFunction создаёт обычный срез, добавляет value
// и возвращает исходный слайс вместе с получившимся окном.
// TODO: поведение должно зависеть от доступной capacity: append может
// изменить исходный массив или перейти на новый.
func AppendWindowInFunction(items []int, low, high, value int) (source, part []int) {
	itemsSlice := items[low:high]
	itemsSlice = append(itemsSlice, value)
	return items, itemsSlice
}

// 08. AppendLimitedWindowInFunction создаёт окно с capacity, ограниченной high,
// затем добавляет value и возвращает исходный слайс и новое окно.
// TODO: добавление не должно перезаписывать элемент исходного слайса за high.
func AppendLimitedWindowInFunction(items []int, low, high, value int) (source, part []int) {
	itemsSlice := items[low:high:high]
	itemsSlice = append(itemsSlice, value)
	return items, itemsSlice
}

// 09. Clone возвращает независимую копию значений.
// TODO: изменение результата не должно менять items. Для nil верните nil.
func Clone(items []int) []int {
	if items == nil || len(items) == 0 {
		return items
	}
	clone := make([]int, len(items))
	copy(clone, items)
	return clone
}

// 10. ChangeClone создаёт независимую копию, меняет в ней элемент index
// и возвращает исходный слайс и копию.
// TODO: в тестах index корректен.
func ChangeClone(items []int, index, value int) (source, clone []int) {
	clone = make([]int, len(items))
	copy(clone, items)
	clone[index] = value
	return items, clone
}

// 11. AppendOne добавляет один элемент и возвращает результат append.
// TODO: сохраните исходный порядок элементов.
func AppendOne(items []int, value int) []int {
	return append(items, value)
}

// 12. AppendMany добавляет все values в исходном порядке.
// TODO: корректно обработайте пустой список добавляемых значений.
func AppendMany(items []int, values ...int) []int {
	return append(items, values...)
}

// 13. CanAppendWithoutGrow сообщает, хватает ли текущей capacity для extra
// новых элементов.
// TODO: отрицательное extra считается некорректным и даёт false.
func CanAppendWithoutGrow(items []int, extra int) bool {
	if extra < 0 {
		return false
	}
	return extra <= cap(items)-len(items)
}

// 14. SliceKind классифицирует слайс.
// TODO: nil -> "nil", ненулевой слайс с len=0 -> "empty",
// остальные значения -> "filled".
func SliceKind(items []int) string {
	switch {
	case items == nil:
		return "nil"
	case len(items) == 0:
		return "empty"
	default:
		return "filled"
	}
}

// 15. AppendIndependent возвращает независимый результат items+values.
// TODO: исходный слайс и его базовый массив не должны измениться,
// даже если в исходной capacity есть свободное место.
func AppendIndependent(items []int, values ...int) []int {
	clone := make([]int, len(items))
	copy(clone, items)
	clone = append(clone, values...)
	return clone
}
