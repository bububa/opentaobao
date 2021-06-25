package product

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
