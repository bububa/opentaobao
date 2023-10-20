package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPromotionListAPIRequest 获取促销规则列表 API请求
// alibaba.alsc.crm.promotion.list
//
// 获取品牌的促销规则列表
type AlibabaAlscCrmPromotionListAPIRequest struct {
	model.Params
	// 获取促销规则请求参数
	_promotionFacadeOpenReq *PromotionFacadeOpenReq
}

// NewAlibabaAlscCrmPromotionListRequest 初始化AlibabaAlscCrmPromotionListAPIRequest对象
func NewAlibabaAlscCrmPromotionListRequest() *AlibabaAlscCrmPromotionListAPIRequest {
	return &AlibabaAlscCrmPromotionListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmPromotionListAPIRequest) Reset() {
	r._promotionFacadeOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPromotionListAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.promotion.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPromotionListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmPromotionListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionFacadeOpenReq is PromotionFacadeOpenReq Setter
// 获取促销规则请求参数
func (r *AlibabaAlscCrmPromotionListAPIRequest) SetPromotionFacadeOpenReq(_promotionFacadeOpenReq *PromotionFacadeOpenReq) error {
	r._promotionFacadeOpenReq = _promotionFacadeOpenReq
	r.Set("promotion_facade_open_req", _promotionFacadeOpenReq)
	return nil
}

// GetPromotionFacadeOpenReq PromotionFacadeOpenReq Getter
func (r AlibabaAlscCrmPromotionListAPIRequest) GetPromotionFacadeOpenReq() *PromotionFacadeOpenReq {
	return r._promotionFacadeOpenReq
}

var poolAlibabaAlscCrmPromotionListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmPromotionListRequest()
	},
}

// GetAlibabaAlscCrmPromotionListRequest 从 sync.Pool 获取 AlibabaAlscCrmPromotionListAPIRequest
func GetAlibabaAlscCrmPromotionListAPIRequest() *AlibabaAlscCrmPromotionListAPIRequest {
	return poolAlibabaAlscCrmPromotionListAPIRequest.Get().(*AlibabaAlscCrmPromotionListAPIRequest)
}

// ReleaseAlibabaAlscCrmPromotionListAPIRequest 将 AlibabaAlscCrmPromotionListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmPromotionListAPIRequest(v *AlibabaAlscCrmPromotionListAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmPromotionListAPIRequest.Put(v)
}
