package wdklogistics

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) Reset() {
	r._stationCode = ""
	r._carNum = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.logistics.pus.pickup.cararrived"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkLogisticsPusPickupCararrivedAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkLogisticsPusPickupCararrivedRequest()
	},
}

// GetAlibabaWdkLogisticsPusPickupCararrivedRequest 从 sync.Pool 获取 AlibabaWdkLogisticsPusPickupCararrivedAPIRequest
func GetAlibabaWdkLogisticsPusPickupCararrivedAPIRequest() *AlibabaWdkLogisticsPusPickupCararrivedAPIRequest {
	return poolAlibabaWdkLogisticsPusPickupCararrivedAPIRequest.Get().(*AlibabaWdkLogisticsPusPickupCararrivedAPIRequest)
}

// ReleaseAlibabaWdkLogisticsPusPickupCararrivedAPIRequest 将 AlibabaWdkLogisticsPusPickupCararrivedAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkLogisticsPusPickupCararrivedAPIRequest(v *AlibabaWdkLogisticsPusPickupCararrivedAPIRequest) {
	v.Reset()
	poolAlibabaWdkLogisticsPusPickupCararrivedAPIRequest.Put(v)
}
