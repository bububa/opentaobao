package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest 应用风险详细信息批量查询接口 API请求
// alibaba.security.jaq.app.riskdetailbatch.get
//
// 用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
type AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
	// 本地化语言信息,用于指定返回结果内容所使用的语言(默认为zh_CN,目前仅支持zh_CN)
	_locale *Locale
}

// NewAlibabaSecurityJaqAppRiskdetailbatchGetRequest 初始化AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest对象
func NewAlibabaSecurityJaqAppRiskdetailbatchGetRequest() *AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest {
	return &AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) Reset() {
	r._itemId = ""
	r._locale = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.riskdetailbatch.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 任务唯一标识
func (r *AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) GetItemId() string {
	return r._itemId
}

// SetLocale is Locale Setter
// 本地化语言信息,用于指定返回结果内容所使用的语言(默认为zh_CN,目前仅支持zh_CN)
func (r *AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) SetLocale(_locale *Locale) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) GetLocale() *Locale {
	return r._locale
}

var poolAlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqAppRiskdetailbatchGetRequest()
	},
}

// GetAlibabaSecurityJaqAppRiskdetailbatchGetRequest 从 sync.Pool 获取 AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest
func GetAlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest() *AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest {
	return poolAlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest.Get().(*AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest)
}

// ReleaseAlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest 将 AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest(v *AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest.Put(v)
}
