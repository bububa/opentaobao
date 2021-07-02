package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppShieldresultGetAPIRequest 用户查询加固结果 API请求
// alibaba.security.jaq.app.shieldresult.get
//
// 用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
type AlibabaSecurityJaqAppShieldresultGetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
}

// NewAlibabaSecurityJaqAppShieldresultGetRequest 初始化AlibabaSecurityJaqAppShieldresultGetAPIRequest对象
func NewAlibabaSecurityJaqAppShieldresultGetRequest() *AlibabaSecurityJaqAppShieldresultGetAPIRequest {
	return &AlibabaSecurityJaqAppShieldresultGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppShieldresultGetAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.shieldresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppShieldresultGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetItemId is ItemId Setter
// 任务唯一标识
func (r *AlibabaSecurityJaqAppShieldresultGetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaSecurityJaqAppShieldresultGetAPIRequest) GetItemId() string {
	return r._itemId
}
