package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRecommendMixcardinfoGetAPIRequest 快应用混合卡片信息 API请求
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
type AlibabaAlihealthRecommendMixcardinfoGetAPIRequest struct {
	model.Params
	// 请求入参
	_quickAppRequest *QuickAppRequest
}

// NewAlibabaAlihealthRecommendMixcardinfoGetRequest 初始化AlibabaAlihealthRecommendMixcardinfoGetAPIRequest对象
func NewAlibabaAlihealthRecommendMixcardinfoGetRequest() *AlibabaAlihealthRecommendMixcardinfoGetAPIRequest {
	return &AlibabaAlihealthRecommendMixcardinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.recommend.mixcardinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuickAppRequest is QuickAppRequest Setter
// 请求入参
func (r *AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) SetQuickAppRequest(_quickAppRequest *QuickAppRequest) error {
	r._quickAppRequest = _quickAppRequest
	r.Set("quick_app_request", _quickAppRequest)
	return nil
}

// GetQuickAppRequest QuickAppRequest Getter
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetQuickAppRequest() *QuickAppRequest {
	return r._quickAppRequest
}
