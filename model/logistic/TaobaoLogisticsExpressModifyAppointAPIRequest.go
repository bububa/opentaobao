package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressmodifyappointAPIRequest 快递改约api API请求
// taobao.logistics.express.modify.appoint
//
// 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
type TaobaologisticsexpressmodifyappointAPIRequest struct {
	model.Params
	// 改约请求对象
	_expressModifyAppointTopRequest *ExpressModifyAppointTopRequestDto
}

// NewTaobaologisticsexpressmodifyappointRequest 初始化TaobaologisticsexpressmodifyappointAPIRequest对象
func NewTaobaologisticsexpressmodifyappointRequest() *TaobaologisticsexpressmodifyappointAPIRequest {
	return &TaobaologisticsexpressmodifyappointAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressmodifyappointAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.modify.appoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressmodifyappointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressmodifyappointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExpressModifyAppointTopRequest is ExpressModifyAppointTopRequest Setter
// 改约请求对象
func (r *TaobaologisticsexpressmodifyappointAPIRequest) SetExpressModifyAppointTopRequest(_expressModifyAppointTopRequest *ExpressModifyAppointTopRequestDto) error {
	r._expressModifyAppointTopRequest = _expressModifyAppointTopRequest
	r.Set("express_modify_appoint_top_request", _expressModifyAppointTopRequest)
	return nil
}

// GetExpressModifyAppointTopRequest ExpressModifyAppointTopRequest Getter
func (r TaobaologisticsexpressmodifyappointAPIRequest) GetExpressModifyAppointTopRequest() *ExpressModifyAppointTopRequestDto {
	return r._expressModifyAppointTopRequest
}
