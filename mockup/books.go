package mockup

type book struct {
	BookId    string `json:"book_id"`
	BookName  string `json:"book_name"`
	BookPrice int    `json:"book_price"`
}

var books = []book{
	{BookId: "1", BookName: "Spider Man", BookPrice: 100},
	{BookId: "2", BookName: "Marvel", BookPrice: 200},
	{BookId: "3", BookName: "Travel", BookPrice: 300},
}

func GetBooks() []book {
	return books
}
