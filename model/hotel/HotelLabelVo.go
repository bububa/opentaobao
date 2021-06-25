package hotel

// HotelLabelVo 
type HotelLabelVo struct {

    // 酒店标签id
    Id   int64 `json:"id,omitempty"`

    // 酒店标签文案
    Text   string `json:"text,omitempty"`

    // 酒店标签颜色
    Color   string `json:"color,omitempty"`

    // desc
    Desc   string `json:"desc,omitempty"`

    // icon
    Icon   string `json:"icon,omitempty"`

    // index
    Index   int64 `json:"index,omitempty"`

    // pos
    Pos   int64 `json:"pos,omitempty"`

    // refFieldName
    RefFieldName   string `json:"ref_field_name,omitempty"`

    // reqFieldName
    ReqFieldName   string `json:"req_field_name,omitempty"`

    // show
    Show   bool `json:"show,omitempty"`

    // value
    Value   int64 `json:"value,omitempty"`

}
