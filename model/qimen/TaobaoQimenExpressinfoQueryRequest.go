package qimen

// TaobaoQimenExpressinfoQueryRequest 
type TaobaoQimenExpressinfoQueryRequest struct {
    // 奇门仓储字段
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 奇门仓储字段
    ExpressCode   string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
    // 扩展属性
    ExtendProps   *TaobaoQimenExpressinfoQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
