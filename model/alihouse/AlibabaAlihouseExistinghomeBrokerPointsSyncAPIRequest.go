package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest 经纪人积分同步 API请求
// alibaba.alihouse.existinghome.broker.points.sync
//
// 经纪人积分
type AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest struct {
	model.Params
	// 经纪人积分列表
	_brokerPointsList *SyncBrokerPointsDto
}

// NewAlibabaAlihouseExistinghomeBrokerPointsSyncRequest 初始化AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeBrokerPointsSyncRequest() *AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.points.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBrokerPointsList is BrokerPointsList Setter
// 经纪人积分列表
func (r *AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest) SetBrokerPointsList(_brokerPointsList *SyncBrokerPointsDto) error {
	r._brokerPointsList = _brokerPointsList
	r.Set("broker_points_list", _brokerPointsList)
	return nil
}

// GetBrokerPointsList BrokerPointsList Getter
func (r AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest) GetBrokerPointsList() *SyncBrokerPointsDto {
	return r._brokerPointsList
}
