package aeusergrowth

// AliexpressUsergrowthRecommendItemsGetResult 
type AliexpressUsergrowthRecommendItemsGetResult struct {

    // Result itemList,The product are located at the top,maybe null when success = false
    
    DataList   []AliexpressUsergrowthRecommendItemsGetData `json:"data_list,omitempty" xml:"data_list,omitempty"`
    

    // success is used to determine whether invoke service success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // other extend message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // Code is used to determine whether the result is correct
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

}
