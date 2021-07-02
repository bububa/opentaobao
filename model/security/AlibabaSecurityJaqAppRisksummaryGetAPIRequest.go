package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRisksummaryGetAPIRequest 应用风险概要信息查询接口 API请求
// alibaba.security.jaq.app.risksummary.get
//
// 用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险概要信息，本接口不返回风险详细信息
type AlibabaSecurityJaqAppRisksummaryGetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
}

// NewAlibabaSecurityJaqAppRisksummaryGetRequest 初始化AlibabaSecurityJaqAppRisksummaryGetAPIRequest对象
func NewAlibabaSecurityJaqAppRisksummaryGetRequest() *AlibabaSecurityJaqAppRisksummaryGetAPIRequest {
	return &AlibabaSecurityJaqAppRisksummaryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRisksummaryGetAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.risksummary.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRisksummaryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 任务唯一标识
func (r *AlibabaSecurityJaqAppRisksummaryGetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlibabaSecurityJaqAppRisksummaryGetAPIRequest) GetItemId() string {
	return r._itemId
}
