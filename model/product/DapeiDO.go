package product

// DapeiDO 
type DapeiDO struct {
    // id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // title
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // desc
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // url
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // items
    Items   []DapeiTemplateItem `json:"items,omitempty" xml:"items>dapei_template_item,omitempty"`
}
