package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthrecommendcardinfogetAPIRequest 快应用卡片信息 API请求
// alibaba.alihealth.recommend.cardinfo.get
//
// 快应用卡片信息
type AlibabaalihealthrecommendcardinfogetAPIRequest struct {
	model.Params
	// 请求入参
	_cardRequest *QuickAppRequest
}

// NewAlibabaalihealthrecommendcardinfogetRequest 初始化AlibabaalihealthrecommendcardinfogetAPIRequest对象
func NewAlibabaalihealthrecommendcardinfogetRequest() *AlibabaalihealthrecommendcardinfogetAPIRequest {
	return &AlibabaalihealthrecommendcardinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthrecommendcardinfogetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.recommend.cardinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthrecommendcardinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthrecommendcardinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardRequest is CardRequest Setter
// 请求入参
func (r *AlibabaalihealthrecommendcardinfogetAPIRequest) SetCardRequest(_cardRequest *QuickAppRequest) error {
	r._cardRequest = _cardRequest
	r.Set("card_request", _cardRequest)
	return nil
}

// GetCardRequest CardRequest Getter
func (r AlibabaalihealthrecommendcardinfogetAPIRequest) GetCardRequest() *QuickAppRequest {
	return r._cardRequest
}
