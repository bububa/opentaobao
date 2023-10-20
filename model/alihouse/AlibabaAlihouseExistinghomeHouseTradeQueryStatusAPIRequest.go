package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest 查询房源状态接口 API请求
// alibaba.alihouse.existinghome.house.trade.query.status
//
// 查询房源状态接口
type AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest struct {
	model.Params
	// 外部小区id
	_communityOuterId string
	// 外部房源ID
	_outerId string
}

// NewAlibabaAlihouseExistinghomeHouseTradeQueryStatusRequest 初始化AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseTradeQueryStatusRequest() *AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.trade.query.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommunityOuterId is CommunityOuterId Setter
// 外部小区id
func (r *AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest) SetCommunityOuterId(_communityOuterId string) error {
	r._communityOuterId = _communityOuterId
	r.Set("community_outer_id", _communityOuterId)
	return nil
}

// GetCommunityOuterId CommunityOuterId Getter
func (r AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest) GetCommunityOuterId() string {
	return r._communityOuterId
}

// SetOuterId is OuterId Setter
// 外部房源ID
func (r *AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseExistinghomeHouseTradeQueryStatusAPIRequest) GetOuterId() string {
	return r._outerId
}
