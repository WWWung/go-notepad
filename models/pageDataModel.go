package models

//PageDataModel ..
type PageDataModel struct {
	PageIndex int         `json:"pageIndex"`
	RowsIndex int         `json:"rowsIndex"`
	PageCount int         `json:"pageCount"`
	Total     int         `json:"total"`
	Rows      interface{} `json:"rows"`
}

//PagingData ..
func PagingData(pageIndex int, rowsIndex int, pageCount int, total int, rows interface{}) PageDataModel {
	return PageDataModel{pageIndex, rowsIndex, pageCount, total, rows}
}
