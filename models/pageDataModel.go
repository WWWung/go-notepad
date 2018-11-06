package models

//PageDataModel ..
type PageDataModel struct {
	PageIndex  int         `json:"pageIndex"`
	RowsInPage int         `json:"rowsInPage"`
	PageCount  int         `json:"pageCount"`
	Total      int         `json:"total"`
	Rows       interface{} `json:"rows"`
}

//PagingData ..
func PagingData(pageIndex int, rowsInPage int, pageCount int, total int, rows interface{}) PageDataModel {
	return PageDataModel{pageIndex, rowsInPage, pageCount, total, rows}
}
