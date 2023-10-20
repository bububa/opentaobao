package wdklogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdklogisticspuspickupcararrivedAPIRequest 自提业务-车辆到达上报车牌号 API请求
// alibaba.wdk.logistics.pus.pickup.cararrived
//
// 自提业务-汽车自提,车辆到达上报车牌号
type AlibabawdklogisticspuspickupcararrivedAPIRequest struct {
	model.Params
	// 自提点
	_stationCode string
	// 车牌号
	_carNum string
}

// NewAlibabawdklogisticspuspickupcararrivedRequest 初始化AlibabawdklogisticspuspickupcararrivedAPIRequest对象
func NewAlibabawdklogisticspuspickupcararrivedRequest() *AlibabawdklogisticspuspickupcararrivedAPIRequest {
	return &AlibabawdklogisticspuspickupcararrivedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdklogisticspuspickupcararrivedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.logistics.pus.pickup.cararrived"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdklogisticspuspickupcararrivedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdklogisticspuspickupcararrivedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStationCode is StationCode Setter
// 自提点
func (r *AlibabawdklogisticspuspickupcararrivedAPIRequest) SetStationCode(_stationCode string) error {
	r._stationCode = _stationCode
	r.Set("station_code", _stationCode)
	return nil
}

// GetStationCode StationCode Getter
func (r AlibabawdklogisticspuspickupcararrivedAPIRequest) GetStationCode() string {
	return r._stationCode
}

// SetCarNum is CarNum Setter
// 车牌号
func (r *AlibabawdklogisticspuspickupcararrivedAPIRequest) SetCarNum(_carNum string) error {
	r._carNum = _carNum
	r.Set("car_num", _carNum)
	return nil
}

// GetCarNum CarNum Getter
func (r AlibabawdklogisticspuspickupcararrivedAPIRequest) GetCarNum() string {
	return r._carNum
}
