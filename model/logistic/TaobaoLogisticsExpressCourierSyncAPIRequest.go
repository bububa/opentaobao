package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresscouriersyncAPIRequest 快递公司同步小件员信息 API请求
// taobao.logistics.express.courier.sync
//
// 快递公司同步小件员信息
type TaobaologisticsexpresscouriersyncAPIRequest struct {
	model.Params
	// 小件员信息
	_tmsCourierRequest *TmsCourierRequest
}

// NewTaobaologisticsexpresscouriersyncRequest 初始化TaobaologisticsexpresscouriersyncAPIRequest对象
func NewTaobaologisticsexpresscouriersyncRequest() *TaobaologisticsexpresscouriersyncAPIRequest {
	return &TaobaologisticsexpresscouriersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpresscouriersyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.courier.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpresscouriersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpresscouriersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsCourierRequest is TmsCourierRequest Setter
// 小件员信息
func (r *TaobaologisticsexpresscouriersyncAPIRequest) SetTmsCourierRequest(_tmsCourierRequest *TmsCourierRequest) error {
	r._tmsCourierRequest = _tmsCourierRequest
	r.Set("tms_courier_request", _tmsCourierRequest)
	return nil
}

// GetTmsCourierRequest TmsCourierRequest Getter
func (r TaobaologisticsexpresscouriersyncAPIRequest) GetTmsCourierRequest() *TmsCourierRequest {
	return r._tmsCourierRequest
}
