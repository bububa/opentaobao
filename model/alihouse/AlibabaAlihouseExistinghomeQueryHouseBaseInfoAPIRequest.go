package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest 查询房源基本信息 API请求
// alibaba.alihouse.existinghome.query.house.base.info
//
// 查询房源基本信息
type AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest struct {
	model.Params
	// 1
	_communityOuterId string
	// 1
	_outerId string
}

// NewAlibabaalihouseexistinghomequeryhousebaseinfoRequest 初始化AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest对象
func NewAlibabaalihouseexistinghomequeryhousebaseinfoRequest() *AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest {
	return &AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.query.house.base.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommunityOuterId is CommunityOuterId Setter
// 1
func (r *AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest) SetCommunityOuterId(_communityOuterId string) error {
	r._communityOuterId = _communityOuterId
	r.Set("community_outer_id", _communityOuterId)
	return nil
}

// GetCommunityOuterId CommunityOuterId Getter
func (r AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest) GetCommunityOuterId() string {
	return r._communityOuterId
}

// SetOuterId is OuterId Setter
// 1
func (r *AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest) GetOuterId() string {
	return r._outerId
}
