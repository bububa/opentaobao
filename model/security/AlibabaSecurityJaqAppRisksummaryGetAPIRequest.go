package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqapprisksummarygetAPIRequest 应用风险概要信息查询接口 API请求
// alibaba.security.jaq.app.risksummary.get
//
// 用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险概要信息，本接口不返回风险详细信息
type AlibabasecurityjaqapprisksummarygetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
}

// NewAlibabasecurityjaqapprisksummarygetRequest 初始化AlibabasecurityjaqapprisksummarygetAPIRequest对象
func NewAlibabasecurityjaqapprisksummarygetRequest() *AlibabasecurityjaqapprisksummarygetAPIRequest {
	return &AlibabasecurityjaqapprisksummarygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqapprisksummarygetAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.risksummary.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqapprisksummarygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqapprisksummarygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 任务唯一标识
func (r *AlibabasecurityjaqapprisksummarygetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabasecurityjaqapprisksummarygetAPIRequest) GetItemId() string {
	return r._itemId
}
