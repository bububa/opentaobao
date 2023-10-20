package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrokerpointssyncAPIRequest 经纪人积分同步 API请求
// alibaba.alihouse.existinghome.broker.points.sync
//
// 经纪人积分
type AlibabaalihouseexistinghomebrokerpointssyncAPIRequest struct {
	model.Params
	// 经纪人积分列表
	_brokerPointsList *SyncBrokerPointsDto
}

// NewAlibabaalihouseexistinghomebrokerpointssyncRequest 初始化AlibabaalihouseexistinghomebrokerpointssyncAPIRequest对象
func NewAlibabaalihouseexistinghomebrokerpointssyncRequest() *AlibabaalihouseexistinghomebrokerpointssyncAPIRequest {
	return &AlibabaalihouseexistinghomebrokerpointssyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomebrokerpointssyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.points.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomebrokerpointssyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomebrokerpointssyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrokerPointsList is BrokerPointsList Setter
// 经纪人积分列表
func (r *AlibabaalihouseexistinghomebrokerpointssyncAPIRequest) SetBrokerPointsList(_brokerPointsList *SyncBrokerPointsDto) error {
	r._brokerPointsList = _brokerPointsList
	r.Set("broker_points_list", _brokerPointsList)
	return nil
}

// GetBrokerPointsList BrokerPointsList Getter
func (r AlibabaalihouseexistinghomebrokerpointssyncAPIRequest) GetBrokerPointsList() *SyncBrokerPointsDto {
	return r._brokerPointsList
}
