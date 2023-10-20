package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoendpointlockertopordernoticesendqueryAPIRequest 查询订单是否由裹裹发送消息 API请求
// cainiao.endpoint.locker.top.order.noticesend.query
//
// 合作公司查询消息发送的接口，判断是否裹裹发送消息
type CainiaoendpointlockertopordernoticesendqueryAPIRequest struct {
	model.Params
	// 站点id
	_stationId string
	// 运单号
	_mailNo string
	// 收件人手机号
	_getterPhone string
}

// NewCainiaoendpointlockertopordernoticesendqueryRequest 初始化CainiaoendpointlockertopordernoticesendqueryAPIRequest对象
func NewCainiaoendpointlockertopordernoticesendqueryRequest() *CainiaoendpointlockertopordernoticesendqueryAPIRequest {
	return &CainiaoendpointlockertopordernoticesendqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoendpointlockertopordernoticesendqueryAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.noticesend.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoendpointlockertopordernoticesendqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoendpointlockertopordernoticesendqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStationId is StationId Setter
// 站点id
func (r *CainiaoendpointlockertopordernoticesendqueryAPIRequest) SetStationId(_stationId string) error {
	r._stationId = _stationId
	r.Set("station_id", _stationId)
	return nil
}

// GetStationId StationId Getter
func (r CainiaoendpointlockertopordernoticesendqueryAPIRequest) GetStationId() string {
	return r._stationId
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoendpointlockertopordernoticesendqueryAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoendpointlockertopordernoticesendqueryAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetGetterPhone is GetterPhone Setter
// 收件人手机号
func (r *CainiaoendpointlockertopordernoticesendqueryAPIRequest) SetGetterPhone(_getterPhone string) error {
	r._getterPhone = _getterPhone
	r.Set("getter_phone", _getterPhone)
	return nil
}

// GetGetterPhone GetterPhone Getter
func (r CainiaoendpointlockertopordernoticesendqueryAPIRequest) GetGetterPhone() string {
	return r._getterPhone
}
