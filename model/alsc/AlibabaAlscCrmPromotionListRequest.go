package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取促销规则列表 API请求
alibaba.alsc.crm.promotion.list

获取品牌的促销规则列表
*/
type AlibabaAlscCrmPromotionListRequest struct {
    model.Params
    // 获取促销规则请求参数
    _promotionFacadeOpenReq   *PromotionFacadeOpenReq
}

// 初始化AlibabaAlscCrmPromotionListRequest对象
func NewAlibabaAlscCrmPromotionListRequest() *AlibabaAlscCrmPromotionListRequest{
    return &AlibabaAlscCrmPromotionListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPromotionListRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.promotion.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPromotionListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PromotionFacadeOpenReq Setter
// 获取促销规则请求参数
func (r *AlibabaAlscCrmPromotionListRequest) SetPromotionFacadeOpenReq(_promotionFacadeOpenReq *PromotionFacadeOpenReq) error {
    r._promotionFacadeOpenReq = _promotionFacadeOpenReq
    r.Set("promotion_facade_open_req", _promotionFacadeOpenReq)
    return nil
}

// PromotionFacadeOpenReq Getter
func (r AlibabaAlscCrmPromotionListRequest) GetPromotionFacadeOpenReq() *PromotionFacadeOpenReq {
    return r._promotionFacadeOpenReq
}
