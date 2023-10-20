package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthrecommendmixcardinfogetAPIRequest 快应用混合卡片信息 API请求
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
type AlibabaalihealthrecommendmixcardinfogetAPIRequest struct {
	model.Params
	// 请求入参
	_quickAppRequest *QuickAppRequest
}

// NewAlibabaalihealthrecommendmixcardinfogetRequest 初始化AlibabaalihealthrecommendmixcardinfogetAPIRequest对象
func NewAlibabaalihealthrecommendmixcardinfogetRequest() *AlibabaalihealthrecommendmixcardinfogetAPIRequest {
	return &AlibabaalihealthrecommendmixcardinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthrecommendmixcardinfogetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.recommend.mixcardinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthrecommendmixcardinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthrecommendmixcardinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuickAppRequest is QuickAppRequest Setter
// 请求入参
func (r *AlibabaalihealthrecommendmixcardinfogetAPIRequest) SetQuickAppRequest(_quickAppRequest *QuickAppRequest) error {
	r._quickAppRequest = _quickAppRequest
	r.Set("quick_app_request", _quickAppRequest)
	return nil
}

// GetQuickAppRequest QuickAppRequest Getter
func (r AlibabaalihealthrecommendmixcardinfogetAPIRequest) GetQuickAppRequest() *QuickAppRequest {
	return r._quickAppRequest
}
