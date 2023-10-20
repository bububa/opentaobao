package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresspickcodecheckAPIRequest 快递公司取货码校验 API请求
// taobao.logistics.express.pickcode.check
//
// 快递公司取货码校验
type TaobaologisticsexpresspickcodecheckAPIRequest struct {
	model.Params
	// 取件码校验参数
	_tmsPickCodeRequest *TmsPickCodeRequest
}

// NewTaobaologisticsexpresspickcodecheckRequest 初始化TaobaologisticsexpresspickcodecheckAPIRequest对象
func NewTaobaologisticsexpresspickcodecheckRequest() *TaobaologisticsexpresspickcodecheckAPIRequest {
	return &TaobaologisticsexpresspickcodecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpresspickcodecheckAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.pickcode.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpresspickcodecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpresspickcodecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsPickCodeRequest is TmsPickCodeRequest Setter
// 取件码校验参数
func (r *TaobaologisticsexpresspickcodecheckAPIRequest) SetTmsPickCodeRequest(_tmsPickCodeRequest *TmsPickCodeRequest) error {
	r._tmsPickCodeRequest = _tmsPickCodeRequest
	r.Set("tms_pick_code_request", _tmsPickCodeRequest)
	return nil
}

// GetTmsPickCodeRequest TmsPickCodeRequest Getter
func (r TaobaologisticsexpresspickcodecheckAPIRequest) GetTmsPickCodeRequest() *TmsPickCodeRequest {
	return r._tmsPickCodeRequest
}
