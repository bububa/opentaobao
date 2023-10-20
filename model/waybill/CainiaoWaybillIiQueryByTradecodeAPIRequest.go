package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaowaybilliiquerybytradecodeAPIRequest 通过订单号查询电子面单通接口 API请求
// cainiao.waybill.ii.query.by.tradecode
//
// 通过订单号查看面单的信息
type CainiaowaybilliiquerybytradecodeAPIRequest struct {
	model.Params
	// 订单号列表
	_paramList []WaybillDetailQueryByBizSubCodeRequest
}

// NewCainiaowaybilliiquerybytradecodeRequest 初始化CainiaowaybilliiquerybytradecodeAPIRequest对象
func NewCainiaowaybilliiquerybytradecodeRequest() *CainiaowaybilliiquerybytradecodeAPIRequest {
	return &CainiaowaybilliiquerybytradecodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaowaybilliiquerybytradecodeAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.query.by.tradecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaowaybilliiquerybytradecodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaowaybilliiquerybytradecodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 订单号列表
func (r *CainiaowaybilliiquerybytradecodeAPIRequest) SetParamList(_paramList []WaybillDetailQueryByBizSubCodeRequest) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r CainiaowaybilliiquerybytradecodeAPIRequest) GetParamList() []WaybillDetailQueryByBizSubCodeRequest {
	return r._paramList
}
