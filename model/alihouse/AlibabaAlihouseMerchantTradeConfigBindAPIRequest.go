package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousemerchanttradeconfigbindAPIRequest 交易场景绑定 API请求
// alibaba.alihouse.merchant.trade.config.bind
//
// 交易场景绑定
type AlibabaalihousemerchanttradeconfigbindAPIRequest struct {
	model.Params
	// 进件资料对象
	_cis *TradeSceneAddInfoDto
}

// NewAlibabaalihousemerchanttradeconfigbindRequest 初始化AlibabaalihousemerchanttradeconfigbindAPIRequest对象
func NewAlibabaalihousemerchanttradeconfigbindRequest() *AlibabaalihousemerchanttradeconfigbindAPIRequest {
	return &AlibabaalihousemerchanttradeconfigbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousemerchanttradeconfigbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.merchant.trade.config.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousemerchanttradeconfigbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousemerchanttradeconfigbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCis is Cis Setter
// 进件资料对象
func (r *AlibabaalihousemerchanttradeconfigbindAPIRequest) SetCis(_cis *TradeSceneAddInfoDto) error {
	r._cis = _cis
	r.Set("cis", _cis)
	return nil
}

// GetCis Cis Getter
func (r AlibabaalihousemerchanttradeconfigbindAPIRequest) GetCis() *TradeSceneAddInfoDto {
	return r._cis
}
