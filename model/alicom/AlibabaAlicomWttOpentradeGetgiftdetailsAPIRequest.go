package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest 存送业务查询奖品信息 API请求
// alibaba.alicom.wtt.opentrade.getgiftdetails
//
// 话费宝充值送查询奖品信息
type AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest struct {
	model.Params
	// 活动ID
	_activityId string
}

// NewAlibabaAlicomWttOpentradeGetgiftdetailsRequest 初始化AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest对象
func NewAlibabaAlicomWttOpentradeGetgiftdetailsRequest() *AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest {
	return &AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.wtt.opentrade.getgiftdetails"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetActivityId is ActivityId Setter
// 活动ID
func (r *AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest) GetActivityId() string {
	return r._activityId
}
