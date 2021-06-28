package wdk

// HeartBeatBo 
/* model for simplify = false
type HeartBeatBo struct {

    // 当前版本信息
    
    VersionId   int64 `json:"version_id,omitempty"`
    

    // MARKET-营销，ITEM-商品
    
    BizCode   string `json:"biz_code,omitempty"`
    

}
*/

// HeartBeatBo 
type HeartBeatBo struct {

    // 当前版本信息
    VersionId   int64 `json:"version_id,omitempty"`

    // MARKET-营销，ITEM-商品
    BizCode   string `json:"biz_code,omitempty"`

}
