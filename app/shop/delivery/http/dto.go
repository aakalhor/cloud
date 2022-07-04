package http

type CreateShop struct {
}

type RegisterItem struct {
	CategoryName string `json:"categoryName"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Price        int    `json:"price"`
	Count        int    `json:"count"`
}

type RegisterCategory struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type AddItem struct {
	Id    string `json:"id"`
	Count int    `json:"count"`
}
