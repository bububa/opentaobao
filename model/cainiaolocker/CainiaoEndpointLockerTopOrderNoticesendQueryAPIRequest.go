package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest
查询订单是否由裹裹发送消息 API请求
cainiao.endpoint.locker.top.order.noticesend.query

合作公司查询消息发送的接口，判断是否裹裹发送消息 */
type CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest struct {
	model.Params
	// 站点id
	_stationId string
	// 收件人手机号
	_getterPhone string
	// 运单号
	_mailNo string
}

// NewCainiaoEndpointLockerTopOrderNoticesendQueryRequest 初始化CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest对象
func NewCainiaoEndpointLockerTopOrderNoticesendQueryRequest() *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest {
	return &CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.noticesend.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StationId Setter
// 站点id
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) SetStationId(_stationId string) error {
	r._stationId = _stationId
	r.Set("station_id", _stationId)
	return nil
}

// Get StationId Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetStationId() string {
	return r._stationId
}

// Set is GetterPhone Setter
// 收件人手机号
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) SetGetterPhone(_getterPhone string) error {
	r._getterPhone = _getterPhone
	r.Set("getter_phone", _getterPhone)
	return nil
}

// Get GetterPhone Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetGetterPhone() string {
	return r._getterPhone
}

// Set is MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// Get MailNo Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetMailNo() string {
	return r._mailNo
}
