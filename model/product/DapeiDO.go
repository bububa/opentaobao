package product

// DapeiDO 
/* model for simplify = false
type DapeiDO struct {

    // id
    
    Id   int64 `json:"id,omitempty"`
    

    // title
    
    Title   string `json:"title,omitempty"`
    

    // desc
    
    Desc   string `json:"desc,omitempty"`
    

    // url
    
    Url   string `json:"url,omitempty"`
    

    // items
    
    Items  struct {
        DapeiTemplateItem  []DapeiTemplateItem `json:"dapei_template_item,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

// DapeiDO 
type DapeiDO struct {

    // id
    Id   int64 `json:"id,omitempty"`

    // title
    Title   string `json:"title,omitempty"`

    // desc
    Desc   string `json:"desc,omitempty"`

    // url
    Url   string `json:"url,omitempty"`

    // items
    Items   []DapeiTemplateItem `json:"items,omitempty"`

}
