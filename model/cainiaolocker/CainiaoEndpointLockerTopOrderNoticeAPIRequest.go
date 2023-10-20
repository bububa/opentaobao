package cainiaolocker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopOrderNoticeAPIRequest 手动触发发短信 API请求
// cainiao.endpoint.locker.top.order.notice
//
// 合作公司对订单手动触发短信，有次数限制
type CainiaoEndpointLockerTopOrderNoticeAPIRequest struct {
	model.Params
	// 合作公司唯一订单编号
	_orderCode string
	// 站点ID
	_stationId string
	// 运单号
	_mailNo string
	// 场景编号：0：重发短信，1：催取短信
	_sceneCode int64
}

// NewCainiaoEndpointLockerTopOrderNoticeRequest 初始化CainiaoEndpointLockerTopOrderNoticeAPIRequest对象
func NewCainiaoEndpointLockerTopOrderNoticeRequest() *CainiaoEndpointLockerTopOrderNoticeAPIRequest {
	return &CainiaoEndpointLockerTopOrderNoticeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) Reset() {
	r._orderCode = ""
	r._stationId = ""
	r._mailNo = ""
	r._sceneCode = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 合作公司唯一订单编号
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetStationId is StationId Setter
// 站点ID
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetStationId(_stationId string) error {
	r._stationId = _stationId
	r.Set("station_id", _stationId)
	return nil
}

// GetStationId StationId Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetStationId() string {
	return r._stationId
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetSceneCode is SceneCode Setter
// 场景编号：0：重发短信，1：催取短信
func (r *CainiaoEndpointLockerTopOrderNoticeAPIRequest) SetSceneCode(_sceneCode int64) error {
	r._sceneCode = _sceneCode
	r.Set("scene_code", _sceneCode)
	return nil
}

// GetSceneCode SceneCode Getter
func (r CainiaoEndpointLockerTopOrderNoticeAPIRequest) GetSceneCode() int64 {
	return r._sceneCode
}

var poolCainiaoEndpointLockerTopOrderNoticeAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoEndpointLockerTopOrderNoticeRequest()
	},
}

// GetCainiaoEndpointLockerTopOrderNoticeRequest 从 sync.Pool 获取 CainiaoEndpointLockerTopOrderNoticeAPIRequest
func GetCainiaoEndpointLockerTopOrderNoticeAPIRequest() *CainiaoEndpointLockerTopOrderNoticeAPIRequest {
	return poolCainiaoEndpointLockerTopOrderNoticeAPIRequest.Get().(*CainiaoEndpointLockerTopOrderNoticeAPIRequest)
}

// ReleaseCainiaoEndpointLockerTopOrderNoticeAPIRequest 将 CainiaoEndpointLockerTopOrderNoticeAPIRequest 放入 sync.Pool
func ReleaseCainiaoEndpointLockerTopOrderNoticeAPIRequest(v *CainiaoEndpointLockerTopOrderNoticeAPIRequest) {
	v.Reset()
	poolCainiaoEndpointLockerTopOrderNoticeAPIRequest.Put(v)
}
