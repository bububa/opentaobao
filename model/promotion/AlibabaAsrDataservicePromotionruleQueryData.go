package promotion

// AlibabaAsrDataservicePromotionruleQueryData 
type AlibabaAsrDataservicePromotionruleQueryData struct {

    // poskey
    PosKey   int64 `json:"pos_key,omitempty"`

    // 名称
    Name   string `json:"name,omitempty"`

    // 兑换详情列表
    DetailList   []Detaillist `json:"detail_list,omitempty"`

}