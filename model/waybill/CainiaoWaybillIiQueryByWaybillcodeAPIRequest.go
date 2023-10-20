package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiQueryByWaybillcodeAPIRequest 通过面单号查询面单打印报文 API请求
// cainiao.waybill.ii.query.by.waybillcode
//
// 通过面单号查询面单的打印报文
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
func (r CainiaoWaybillIiQueryByWaybillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillIiQueryByWaybillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
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
