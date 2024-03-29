package mtopopen

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractAllsparkisvDrawAPIRequest) Reset() {
	r._test = ""
	r._ddd = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.allsparkisv.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTest is Test Setter
// ddd
func (r *AlibabaInteractAllsparkisvDrawAPIRequest) SetTest(_test string) error {
	r._test = _test
	r.Set("test", _test)
	return nil
}

// GetTest Test Getter
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetTest() string {
	return r._test
}

// SetDdd is Ddd Setter
// dd
func (r *AlibabaInteractAllsparkisvDrawAPIRequest) SetDdd(_ddd string) error {
	r._ddd = _ddd
	r.Set("ddd", _ddd)
	return nil
}

// GetDdd Ddd Getter
func (r AlibabaInteractAllsparkisvDrawAPIRequest) GetDdd() string {
	return r._ddd
}

var poolAlibabaInteractAllsparkisvDrawAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractAllsparkisvDrawRequest()
	},
}

// GetAlibabaInteractAllsparkisvDrawRequest 从 sync.Pool 获取 AlibabaInteractAllsparkisvDrawAPIRequest
func GetAlibabaInteractAllsparkisvDrawAPIRequest() *AlibabaInteractAllsparkisvDrawAPIRequest {
	return poolAlibabaInteractAllsparkisvDrawAPIRequest.Get().(*AlibabaInteractAllsparkisvDrawAPIRequest)
}

// ReleaseAlibabaInteractAllsparkisvDrawAPIRequest 将 AlibabaInteractAllsparkisvDrawAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractAllsparkisvDrawAPIRequest(v *AlibabaInteractAllsparkisvDrawAPIRequest) {
	v.Reset()
	poolAlibabaInteractAllsparkisvDrawAPIRequest.Put(v)
}
