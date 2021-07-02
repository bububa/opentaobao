package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiQueryByWaybillcodeAPIRequest 通过面单号查询面单信息 API请求
// cainiao.waybill.ii.query.by.waybillcode
//
// 通过面单号查看面单号的当前状态，如签收、发货、失效等。
type CainiaoWaybillIiQueryByWaybillcodeAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramList []WaybillDetailQueryByWaybillCodeRequest
}

// NewCainiaoWaybillIiQueryByWaybillcodeRequest 初始化CainiaoWaybillIiQueryByWaybillcodeAPIRequest对象
func NewCainiaoWaybillIiQueryByWaybillcodeRequest() *CainiaoWaybillIiQueryByWaybillcodeAPIRequest {
	return &CainiaoWaybillIiQueryByWaybillcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiQueryByWaybillcodeAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.ii.query.by.waybillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiQueryByWaybillcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamList is ParamList Setter
// 系统自动生成
func (r *CainiaoWaybillIiQueryByWaybillcodeAPIRequest) SetParamList(_paramList []WaybillDetailQueryByWaybillCodeRequest) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r CainiaoWaybillIiQueryByWaybillcodeAPIRequest) GetParamList() []WaybillDetailQueryByWaybillCodeRequest {
	return r._paramList
}
