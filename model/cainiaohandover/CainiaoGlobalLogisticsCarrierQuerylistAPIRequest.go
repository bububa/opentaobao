package cainiaohandover

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalLogisticsCarrierQuerylistAPIRequest 实际承运商查询 API请求
// cainiao.global.logistics.carrier.querylist
//
// 查询出所有的实际承运商
type CainiaoGlobalLogisticsCarrierQuerylistAPIRequest struct {
	model.Params
	// 多语言(暂不支持，保留入参)
	_locale string
}

// NewCainiaoGlobalLogisticsCarrierQuerylistRequest 初始化CainiaoGlobalLogisticsCarrierQuerylistAPIRequest对象
func NewCainiaoGlobalLogisticsCarrierQuerylistRequest() *CainiaoGlobalLogisticsCarrierQuerylistAPIRequest {
	return &CainiaoGlobalLogisticsCarrierQuerylistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) Reset() {
	r._locale = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) GetApiMethodName() string {
	return "cainiao.global.logistics.carrier.querylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocale is Locale Setter
// 多语言(暂不支持，保留入参)
func (r *CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) GetLocale() string {
	return r._locale
}

var poolCainiaoGlobalLogisticsCarrierQuerylistAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalLogisticsCarrierQuerylistRequest()
	},
}

// GetCainiaoGlobalLogisticsCarrierQuerylistRequest 从 sync.Pool 获取 CainiaoGlobalLogisticsCarrierQuerylistAPIRequest
func GetCainiaoGlobalLogisticsCarrierQuerylistAPIRequest() *CainiaoGlobalLogisticsCarrierQuerylistAPIRequest {
	return poolCainiaoGlobalLogisticsCarrierQuerylistAPIRequest.Get().(*CainiaoGlobalLogisticsCarrierQuerylistAPIRequest)
}

// ReleaseCainiaoGlobalLogisticsCarrierQuerylistAPIRequest 将 CainiaoGlobalLogisticsCarrierQuerylistAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalLogisticsCarrierQuerylistAPIRequest(v *CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) {
	v.Reset()
	poolCainiaoGlobalLogisticsCarrierQuerylistAPIRequest.Put(v)
}
