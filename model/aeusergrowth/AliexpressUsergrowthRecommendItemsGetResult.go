package aeusergrowth

// AliexpressusergrowthrecommenditemsgetResult 结构体
type AliexpressusergrowthrecommenditemsgetResult struct {
	// Result itemList,The product are located at the top,maybe null when success = false
	DataList []AliexpressusergrowthrecommenditemsgetData `json:"data_list,omitempty" xml:"data_list>aliexpressusergrowthrecommenditemsget_data,omitempty"`
	// other extend message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// Code is used to determine whether the result is correct
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// success is used to determine whether invoke service success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
