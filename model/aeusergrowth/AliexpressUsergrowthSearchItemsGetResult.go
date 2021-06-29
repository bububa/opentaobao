package aeusergrowth

// AliexpressUsergrowthSearchItemsGetResult 
type AliexpressUsergrowthSearchItemsGetResult struct {

    // Result,The product  are located at the top,maybe null  when success = false
    
    DataList   []AliexpressUsergrowthSearchItemsGetData `json:"data_list,omitempty" xml:"data_list,omitempty"`
    

    // success is used to determine whether invoke service success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // other extend message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // Code is used to determine whether the result is correct
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // extend param
    
    Ext   *Ext `json:"ext,omitempty" xml:"ext,omitempty"`
    

}
