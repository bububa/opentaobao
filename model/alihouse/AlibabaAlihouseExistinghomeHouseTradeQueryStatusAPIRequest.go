package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousetradequerystatusAPIRequest 查询房源状态接口 API请求
// alibaba.alihouse.existinghome.house.trade.query.status
//
// 查询房源状态接口
type AlibabaalihouseexistinghomehousetradequerystatusAPIRequest struct {
	model.Params
	// 外部小区id
	_communityOuterId string
	// 外部房源ID
	_outerId string
}

// NewAlibabaalihouseexistinghomehousetradequerystatusRequest 初始化AlibabaalihouseexistinghomehousetradequerystatusAPIRequest对象
func NewAlibabaalihouseexistinghomehousetradequerystatusRequest() *AlibabaalihouseexistinghomehousetradequerystatusAPIRequest {
	return &AlibabaalihouseexistinghomehousetradequerystatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehousetradequerystatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.trade.query.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehousetradequerystatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehousetradequerystatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommunityOuterId is CommunityOuterId Setter
// 外部小区id
func (r *AlibabaalihouseexistinghomehousetradequerystatusAPIRequest) SetCommunityOuterId(_communityOuterId string) error {
	r._communityOuterId = _communityOuterId
	r.Set("community_outer_id", _communityOuterId)
	return nil
}

// GetCommunityOuterId CommunityOuterId Getter
func (r AlibabaalihouseexistinghomehousetradequerystatusAPIRequest) GetCommunityOuterId() string {
	return r._communityOuterId
}

// SetOuterId is OuterId Setter
// 外部房源ID
func (r *AlibabaalihouseexistinghomehousetradequerystatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihouseexistinghomehousetradequerystatusAPIRequest) GetOuterId() string {
	return r._outerId
}
