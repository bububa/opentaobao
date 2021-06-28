package icbu

// PhotoGroupResult 
/* model for simplify = false
type PhotoGroupResult struct {

    // add操作中表示新增的图片分组，rename操作中表示重命名的分组，delete操作中返回分组信息
    
    PhotobankGroup  *struct {
        PhotobankGroup  *PhotobankGroup `json:"photobank_group,omitempty"`
    } `json:"photobank_group,omitempty"`
    

}
*/

// PhotoGroupResult 
type PhotoGroupResult struct {

    // add操作中表示新增的图片分组，rename操作中表示重命名的分组，delete操作中返回分组信息
    PhotobankGroup   *PhotobankGroup `json:"photobank_group,omitempty"`

}
