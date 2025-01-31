package domain

type Category int

const (
	cities Category = iota + 1
	books
	professions
)

var categoryFilenames = [...]string{"cities", "books", "professions"}
var categoriesInRussian = [...]string{"города", "книги", "профессии"}

const categoriesCount = len(categoriesInRussian)

func FirstCategory() Category {
	return cities
}

func LastCategory() Category {
	return professions
}

func CategoriesCount() int {
	return categoriesCount
}

func (d Category) fileName() string {
	return categoryFilenames[d-1]
}

func CategoryFileNameByIndex(ind int) string {
	for i := cities; i <= professions; i++ {
		if int(i) == ind {
			return i.fileName()
		}
	}

	return ""
}

func CategoryInRussian(d Category) string {
	return categoriesInRussian[d-1]
}
