package cainiaolocker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest 查询订单是否由裹裹发送消息 API请求
// cainiao.endpoint.locker.top.order.noticesend.query
//
// 合作公司查询消息发送的接口，判断是否裹裹发送消息
type CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest struct {
	model.Params
	// 站点id
	_stationId string
	// 运单号
	_mailNo string
	// 收件人手机号
	_getterPhone string
}

// NewCainiaoEndpointLockerTopOrderNoticesendQueryRequest 初始化CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest对象
func NewCainiaoEndpointLockerTopOrderNoticesendQueryRequest() *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest {
	return &CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) Reset() {
	r._stationId = ""
	r._mailNo = ""
	r._getterPhone = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.noticesend.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStationId is StationId Setter
// 站点id
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) SetStationId(_stationId string) error {
	r._stationId = _stationId
	r.Set("station_id", _stationId)
	return nil
}

// GetStationId StationId Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetStationId() string {
	return r._stationId
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetGetterPhone is GetterPhone Setter
// 收件人手机号
func (r *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) SetGetterPhone(_getterPhone string) error {
	r._getterPhone = _getterPhone
	r.Set("getter_phone", _getterPhone)
	return nil
}

// GetGetterPhone GetterPhone Getter
func (r CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) GetGetterPhone() string {
	return r._getterPhone
}

var poolCainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoEndpointLockerTopOrderNoticesendQueryRequest()
	},
}

// GetCainiaoEndpointLockerTopOrderNoticesendQueryRequest 从 sync.Pool 获取 CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest
func GetCainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest() *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest {
	return poolCainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest.Get().(*CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest)
}

// ReleaseCainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest 将 CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest 放入 sync.Pool
func ReleaseCainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest(v *CainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest) {
	v.Reset()
	poolCainiaoEndpointLockerTopOrderNoticesendQueryAPIRequest.Put(v)
}
