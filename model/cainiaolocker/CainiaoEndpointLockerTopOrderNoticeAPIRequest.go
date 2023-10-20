package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoendpointlockertopordernoticeAPIRequest 手动触发发短信 API请求
// cainiao.endpoint.locker.top.order.notice
//
// 合作公司对订单手动触发短信，有次数限制
type CainiaoendpointlockertopordernoticeAPIRequest struct {
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

// NewCainiaoendpointlockertopordernoticeRequest 初始化CainiaoendpointlockertopordernoticeAPIRequest对象
func NewCainiaoendpointlockertopordernoticeRequest() *CainiaoendpointlockertopordernoticeAPIRequest {
	return &CainiaoendpointlockertopordernoticeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoendpointlockertopordernoticeAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.order.notice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoendpointlockertopordernoticeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoendpointlockertopordernoticeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 合作公司唯一订单编号
func (r *CainiaoendpointlockertopordernoticeAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r CainiaoendpointlockertopordernoticeAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetStationId is StationId Setter
// 站点ID
func (r *CainiaoendpointlockertopordernoticeAPIRequest) SetStationId(_stationId string) error {
	r._stationId = _stationId
	r.Set("station_id", _stationId)
	return nil
}

// GetStationId StationId Getter
func (r CainiaoendpointlockertopordernoticeAPIRequest) GetStationId() string {
	return r._stationId
}

// SetMailNo is MailNo Setter
// 运单号
func (r *CainiaoendpointlockertopordernoticeAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r CainiaoendpointlockertopordernoticeAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetSceneCode is SceneCode Setter
// 场景编号：0：重发短信，1：催取短信
func (r *CainiaoendpointlockertopordernoticeAPIRequest) SetSceneCode(_sceneCode int64) error {
	r._sceneCode = _sceneCode
	r.Set("scene_code", _sceneCode)
	return nil
}

// GetSceneCode SceneCode Getter
func (r CainiaoendpointlockertopordernoticeAPIRequest) GetSceneCode() int64 {
	return r._sceneCode
}
