package security

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqAppShieldresultGetAPIRequest) Reset() {
	r._itemId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppShieldresultGetAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.shieldresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppShieldresultGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqAppShieldresultGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaSecurityJaqAppShieldresultGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqAppShieldresultGetRequest()
	},
}

// GetAlibabaSecurityJaqAppShieldresultGetRequest 从 sync.Pool 获取 AlibabaSecurityJaqAppShieldresultGetAPIRequest
func GetAlibabaSecurityJaqAppShieldresultGetAPIRequest() *AlibabaSecurityJaqAppShieldresultGetAPIRequest {
	return poolAlibabaSecurityJaqAppShieldresultGetAPIRequest.Get().(*AlibabaSecurityJaqAppShieldresultGetAPIRequest)
}

// ReleaseAlibabaSecurityJaqAppShieldresultGetAPIRequest 将 AlibabaSecurityJaqAppShieldresultGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqAppShieldresultGetAPIRequest(v *AlibabaSecurityJaqAppShieldresultGetAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqAppShieldresultGetAPIRequest.Put(v)
}
