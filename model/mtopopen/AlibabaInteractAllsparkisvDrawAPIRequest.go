package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractallsparkisvdrawAPIRequest allspark提供抽奖tida接口对应鉴权接口 API请求
// alibaba.interact.allsparkisv.draw
//
// 该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用
type AlibabainteractallsparkisvdrawAPIRequest struct {
	model.Params
	// ddd
	_test string
	// dd
	_ddd string
}

// NewAlibabainteractallsparkisvdrawRequest 初始化AlibabainteractallsparkisvdrawAPIRequest对象
func NewAlibabainteractallsparkisvdrawRequest() *AlibabainteractallsparkisvdrawAPIRequest {
	return &AlibabainteractallsparkisvdrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractallsparkisvdrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.allsparkisv.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractallsparkisvdrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractallsparkisvdrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTest is Test Setter
// ddd
func (r *AlibabainteractallsparkisvdrawAPIRequest) SetTest(_test string) error {
	r._test = _test
	r.Set("test", _test)
	return nil
}

// GetTest Test Getter
func (r AlibabainteractallsparkisvdrawAPIRequest) GetTest() string {
	return r._test
}

// SetDdd is Ddd Setter
// dd
func (r *AlibabainteractallsparkisvdrawAPIRequest) SetDdd(_ddd string) error {
	r._ddd = _ddd
	r.Set("ddd", _ddd)
	return nil
}

// GetDdd Ddd Getter
func (r AlibabainteractallsparkisvdrawAPIRequest) GetDdd() string {
	return r._ddd
}
