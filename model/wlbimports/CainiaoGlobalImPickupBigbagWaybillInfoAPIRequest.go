package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupbigbagwaybillinfoAPIRequest 大包面单查询 API请求
// cainiao.global.im.pickup.bigbag.waybill.info
//
// 大包面单查询
type CainiaoglobalimpickupbigbagwaybillinfoAPIRequest struct {
	model.Params
	// bigbagId和appointmentOrderId必填一个
	_bigbagWaybillRequest *BigbagWaybillRequest
}

// NewCainiaoglobalimpickupbigbagwaybillinfoRequest 初始化CainiaoglobalimpickupbigbagwaybillinfoAPIRequest对象
func NewCainiaoglobalimpickupbigbagwaybillinfoRequest() *CainiaoglobalimpickupbigbagwaybillinfoAPIRequest {
	return &CainiaoglobalimpickupbigbagwaybillinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalimpickupbigbagwaybillinfoAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.waybill.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalimpickupbigbagwaybillinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalimpickupbigbagwaybillinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagWaybillRequest is BigbagWaybillRequest Setter
// bigbagId和appointmentOrderId必填一个
func (r *CainiaoglobalimpickupbigbagwaybillinfoAPIRequest) SetBigbagWaybillRequest(_bigbagWaybillRequest *BigbagWaybillRequest) error {
	r._bigbagWaybillRequest = _bigbagWaybillRequest
	r.Set("bigbag_waybill_request", _bigbagWaybillRequest)
	return nil
}

// GetBigbagWaybillRequest BigbagWaybillRequest Getter
func (r CainiaoglobalimpickupbigbagwaybillinfoAPIRequest) GetBigbagWaybillRequest() *BigbagWaybillRequest {
	return r._bigbagWaybillRequest
}
