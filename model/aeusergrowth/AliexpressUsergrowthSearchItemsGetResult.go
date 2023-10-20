package aeusergrowth

// AliexpressusergrowthsearchitemsgetResult 结构体
type AliexpressusergrowthsearchitemsgetResult struct {
	// Result,The product  are located at the top,maybe null  when success = false
	DataList []AliexpressusergrowthsearchitemsgetData `json:"data_list,omitempty" xml:"data_list>aliexpressusergrowthsearchitemsget_data,omitempty"`
	// other extend message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// Code is used to determine whether the result is correct
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// extend param
	Ext *Ext `json:"ext,omitempty" xml:"ext,omitempty"`
	// success is used to determine whether invoke service success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
