package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiQueryByTradecodeAPIRequest 通过订单号查询电子面单通接口 API请求
// cainiao.waybill.ii.query.by.tradecode
//
// 通过订单号查看面单的信息
type CainiaoWaybillIiQueryByTradecodeAPIRequest struct {
	model.Params
	// 订单号列表
	_paramList []WaybillDetailQueryByBizSubCodeRequest
}

// NewCainiaoWaybillIiQueryByTradecodeRequest 初始化CainiaoWaybillIiQueryByTradecodeAPIRequest对象
func NewCainiaoWaybillIiQueryByTradecodeRequest() *CainiaoWaybillIiQueryByTradecodeAPIRequest {
	return &CainiaoWaybillIiQueryByTradecodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillIiQueryByTradecodeAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiQueryByTradecodeAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.query.by.tradecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiQueryByTradecodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillIiQueryByTradecodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 订单号列表
func (r *CainiaoWaybillIiQueryByTradecodeAPIRequest) SetParamList(_paramList []WaybillDetailQueryByBizSubCodeRequest) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r CainiaoWaybillIiQueryByTradecodeAPIRequest) GetParamList() []WaybillDetailQueryByBizSubCodeRequest {
	return r._paramList
}

var poolCainiaoWaybillIiQueryByTradecodeAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillIiQueryByTradecodeRequest()
	},
}

// GetCainiaoWaybillIiQueryByTradecodeRequest 从 sync.Pool 获取 CainiaoWaybillIiQueryByTradecodeAPIRequest
func GetCainiaoWaybillIiQueryByTradecodeAPIRequest() *CainiaoWaybillIiQueryByTradecodeAPIRequest {
	return poolCainiaoWaybillIiQueryByTradecodeAPIRequest.Get().(*CainiaoWaybillIiQueryByTradecodeAPIRequest)
}

// ReleaseCainiaoWaybillIiQueryByTradecodeAPIRequest 将 CainiaoWaybillIiQueryByTradecodeAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillIiQueryByTradecodeAPIRequest(v *CainiaoWaybillIiQueryByTradecodeAPIRequest) {
	v.Reset()
	poolCainiaoWaybillIiQueryByTradecodeAPIRequest.Put(v)
}
