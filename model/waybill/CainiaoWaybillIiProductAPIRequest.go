package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiProductAPIRequest 商家查询物流商产品类型接口 API请求
// cainiao.waybill.ii.product
//
// 商家可以查询物流商的产品类型和服务能力。
type CainiaoWaybillIiProductAPIRequest struct {
	model.Params
	// 快递公司code
	_cpCode string
}

// NewCainiaoWaybillIiProductRequest 初始化CainiaoWaybillIiProductAPIRequest对象
func NewCainiaoWaybillIiProductRequest() *CainiaoWaybillIiProductAPIRequest {
	return &CainiaoWaybillIiProductAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillIiProductAPIRequest) Reset() {
	r._cpCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiProductAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiProductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillIiProductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpCode is CpCode Setter
// 快递公司code
func (r *CainiaoWaybillIiProductAPIRequest) SetCpCode(_cpCode string) error {
	r._cpCode = _cpCode
	r.Set("cp_code", _cpCode)
	return nil
}

// GetCpCode CpCode Getter
func (r CainiaoWaybillIiProductAPIRequest) GetCpCode() string {
	return r._cpCode
}

var poolCainiaoWaybillIiProductAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillIiProductRequest()
	},
}

// GetCainiaoWaybillIiProductRequest 从 sync.Pool 获取 CainiaoWaybillIiProductAPIRequest
func GetCainiaoWaybillIiProductAPIRequest() *CainiaoWaybillIiProductAPIRequest {
	return poolCainiaoWaybillIiProductAPIRequest.Get().(*CainiaoWaybillIiProductAPIRequest)
}

// ReleaseCainiaoWaybillIiProductAPIRequest 将 CainiaoWaybillIiProductAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillIiProductAPIRequest(v *CainiaoWaybillIiProductAPIRequest) {
	v.Reset()
	poolCainiaoWaybillIiProductAPIRequest.Put(v)
}
