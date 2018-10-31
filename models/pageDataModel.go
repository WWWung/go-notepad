package models

//PageDataModel ..
type PageDataModel struct {
	pageIndex int
	rowsIndex int
	pageCount int
	total     int
	rows      interface{}
}

//PagingData ..
func PagingData(pageIndex int, rowsIndex int, pageCount int, total int, rows interface{}) PageDataModel {
	return PageDataModel{pageIndex, rowsIndex, pageCount, total, rows}
}
