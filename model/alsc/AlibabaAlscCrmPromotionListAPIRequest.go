package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpromotionlistAPIRequest 获取促销规则列表 API请求
// alibaba.alsc.crm.promotion.list
//
// 获取品牌的促销规则列表
type AlibabaalsccrmpromotionlistAPIRequest struct {
	model.Params
	// 获取促销规则请求参数
	_promotionFacadeOpenReq *PromotionFacadeOpenReq
}

// NewAlibabaalsccrmpromotionlistRequest 初始化AlibabaalsccrmpromotionlistAPIRequest对象
func NewAlibabaalsccrmpromotionlistRequest() *AlibabaalsccrmpromotionlistAPIRequest {
	return &AlibabaalsccrmpromotionlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmpromotionlistAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.promotion.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmpromotionlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmpromotionlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionFacadeOpenReq is PromotionFacadeOpenReq Setter
// 获取促销规则请求参数
func (r *AlibabaalsccrmpromotionlistAPIRequest) SetPromotionFacadeOpenReq(_promotionFacadeOpenReq *PromotionFacadeOpenReq) error {
	r._promotionFacadeOpenReq = _promotionFacadeOpenReq
	r.Set("promotion_facade_open_req", _promotionFacadeOpenReq)
	return nil
}

// GetPromotionFacadeOpenReq PromotionFacadeOpenReq Getter
func (r AlibabaalsccrmpromotionlistAPIRequest) GetPromotionFacadeOpenReq() *PromotionFacadeOpenReq {
	return r._promotionFacadeOpenReq
}
