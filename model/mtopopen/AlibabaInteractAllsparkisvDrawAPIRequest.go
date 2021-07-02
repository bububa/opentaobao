package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractAllsparkisvDrawAPIRequest allspark提供抽奖tida接口对应鉴权接口 API请求
// alibaba.interact.allsparkisv.draw
//
// 该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
type AlibabaInteractAllsparkisvDrawAPIRequest struct {
	model.Params
	// ddd
	_test string
	// dd
	_ddd string
}

// NewAlibabaInteractAllsparkisvDrawRequest 初始化AlibabaInteractAllsparkisvDrawAPIRequest对象
func NewAlibabaInteractAllsparkisvDrawRequest() *AlibabaInteractAllsparkisvDrawAPIRequest {
	return &AlibabaInteractAllsparkisvDrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.allsparkisv.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Test Setter
// ddd
func (r *AlibabaInteractAllsparkisvDrawAPIRequest) SetTest(_test string) error {
	r._test = _test
	r.Set("test", _test)
	return nil
}

// Get Test Getter
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetTest() string {
	return r._test
}

// Set is Ddd Setter
// dd
func (r *AlibabaInteractAllsparkisvDrawAPIRequest) SetDdd(_ddd string) error {
	r._ddd = _ddd
	r.Set("ddd", _ddd)
	return nil
}

// Get Ddd Getter
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetDdd() string {
	return r._ddd
}
