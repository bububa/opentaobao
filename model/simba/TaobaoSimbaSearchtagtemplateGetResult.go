package simba

// TaobaoSimbaSearchtagtemplateGetResult 
type TaobaoSimbaSearchtagtemplateGetResult struct {

    // DimDtOs
    DimList   []DimDtOs `json:"dim_list,omitempty"`

    // 人群模版id
    Id   int64 `json:"id,omitempty"`

    // 人群模版名称
    Name   string `json:"name,omitempty"`

    // 人群类型
    Type   int64 `json:"type,omitempty"`

}
