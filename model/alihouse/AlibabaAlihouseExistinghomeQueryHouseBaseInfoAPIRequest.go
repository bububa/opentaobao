package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest 查询房源基本信息 API请求
// alibaba.alihouse.existinghome.query.house.base.info
//
// 查询房源基本信息
type AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest struct {
	model.Params
	// 1
	_communityOuterId string
	// 1
	_outerId string
}

// NewAlibabaAlihouseExistinghomeQueryHouseBaseInfoRequest 初始化AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest对象
func NewAlibabaAlihouseExistinghomeQueryHouseBaseInfoRequest() *AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest {
	return &AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.query.house.base.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommunityOuterId is CommunityOuterId Setter
// 1
func (r *AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest) SetCommunityOuterId(_communityOuterId string) error {
	r._communityOuterId = _communityOuterId
	r.Set("community_outer_id", _communityOuterId)
	return nil
}

// GetCommunityOuterId CommunityOuterId Getter
func (r AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest) GetCommunityOuterId() string {
	return r._communityOuterId
}

// SetOuterId is OuterId Setter
// 1
func (r *AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseExistinghomeQueryHouseBaseInfoAPIRequest) GetOuterId() string {
	return r._outerId
}
