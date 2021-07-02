package wdklogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkLogisticsPusPickupCararrivedAPIRequest 自提业务-车辆到达上报车牌号 API请求
// alibaba.wdk.logistics.pus.pickup.cararrived
//
// 自提业务-汽车自提,车辆到达上报车牌号
type AlibabaWdkLogisticsPusPickupCararrivedAPIRequest struct {
	model.Params
	// 自提点
	_stationCode string
	// 车牌号
	_carNum string
}

// NewAlibabaWdkLogisticsPusPickupCararrivedRequest 初始化AlibabaWdkLogisticsPusPickupCararrivedAPIRequest对象
func NewAlibabaWdkLogisticsPusPickupCararrivedRequest() *AlibabaWdkLogisticsPusPickupCararrivedAPIRequest {
	return &AlibabaWdkLogisticsPusPickupCararrivedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.logistics.pus.pickup.cararrived"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetStationCode is StationCode Setter
// 自提点
func (r *AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) SetStationCode(_stationCode string) error {
	r._stationCode = _stationCode
	r.Set("station_code", _stationCode)
	return nil
}

// GetStationCode StationCode Getter
func (r AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) GetStationCode() string {
	return r._stationCode
}

// SetCarNum is CarNum Setter
// 车牌号
func (r *AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) SetCarNum(_carNum string) error {
	r._carNum = _carNum
	r.Set("car_num", _carNum)
	return nil
}

// GetCarNum CarNum Getter
func (r AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) GetCarNum() string {
	return r._carNum
}
