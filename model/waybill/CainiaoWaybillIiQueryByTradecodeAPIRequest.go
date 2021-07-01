package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiQueryByTradecodeAPIRequest
通过订单号查询电子面单通接口 API请求
cainiao.waybill.ii.query.by.tradecode

通过订单号查看面单的信息 */
type CainiaoWaybillIiQueryByTradecodeAPIRequest struct {
	model.Params
	// 订单号列表
	_paramList []WaybillDetailQueryByBizSubCodeRequest
}

// NewCainiaoWaybillIiQueryByTradecodeRequest 初始化CainiaoWaybillIiQueryByTradecodeAPIRequest对象
func NewCainiaoWaybillIiQueryByTradecodeRequest() *CainiaoWaybillIiQueryByTradecodeAPIRequest {
	return &CainiaoWaybillIiQueryByTradecodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiQueryByTradecodeAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.query.by.tradecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiQueryByTradecodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamList Setter
// 订单号列表
func (r *CainiaoWaybillIiQueryByTradecodeAPIRequest) SetParamList(_paramList []WaybillDetailQueryByBizSubCodeRequest) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// Get ParamList Getter
func (r CainiaoWaybillIiQueryByTradecodeAPIRequest) GetParamList() []WaybillDetailQueryByBizSubCodeRequest {
	return r._paramList
}
