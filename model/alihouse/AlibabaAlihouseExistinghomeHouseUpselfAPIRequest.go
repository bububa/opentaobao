package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehouseupselfAPIRequest 房源上架 API请求
// alibaba.alihouse.existinghome.house.upself
//
// 房源信息上架
type AlibabaalihouseexistinghomehouseupselfAPIRequest struct {
	model.Params
	// 外部小区id
	_communityOuterId string
	// 外部房源id
	_outerId string
	// 外部委托id
	_outerEntrustId string
	// 1-新房，2-小区，3-公寓。与楼盘的枚举对齐，目的是查询小区或者是公寓数据
	_businessType int64
	// 是否异步，0同步，1异步
	_isAsync int64
	// ms级时间戳
	_etcVersion int64
}

// NewAlibabaalihouseexistinghomehouseupselfRequest 初始化AlibabaalihouseexistinghomehouseupselfAPIRequest对象
func NewAlibabaalihouseexistinghomehouseupselfRequest() *AlibabaalihouseexistinghomehouseupselfAPIRequest {
	return &AlibabaalihouseexistinghomehouseupselfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.upself"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommunityOuterId is CommunityOuterId Setter
// 外部小区id
func (r *AlibabaalihouseexistinghomehouseupselfAPIRequest) SetCommunityOuterId(_communityOuterId string) error {
	r._communityOuterId = _communityOuterId
	r.Set("community_outer_id", _communityOuterId)
	return nil
}

// GetCommunityOuterId CommunityOuterId Getter
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetCommunityOuterId() string {
	return r._communityOuterId
}

// SetOuterId is OuterId Setter
// 外部房源id
func (r *AlibabaalihouseexistinghomehouseupselfAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterEntrustId is OuterEntrustId Setter
// 外部委托id
func (r *AlibabaalihouseexistinghomehouseupselfAPIRequest) SetOuterEntrustId(_outerEntrustId string) error {
	r._outerEntrustId = _outerEntrustId
	r.Set("outer_entrust_id", _outerEntrustId)
	return nil
}

// GetOuterEntrustId OuterEntrustId Getter
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetOuterEntrustId() string {
	return r._outerEntrustId
}

// SetBusinessType is BusinessType Setter
// 1-新房，2-小区，3-公寓。与楼盘的枚举对齐，目的是查询小区或者是公寓数据
func (r *AlibabaalihouseexistinghomehouseupselfAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetBusinessType() int64 {
	return r._businessType
}

// SetIsAsync is IsAsync Setter
// 是否异步，0同步，1异步
func (r *AlibabaalihouseexistinghomehouseupselfAPIRequest) SetIsAsync(_isAsync int64) error {
	r._isAsync = _isAsync
	r.Set("is_async", _isAsync)
	return nil
}

// GetIsAsync IsAsync Getter
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetIsAsync() int64 {
	return r._isAsync
}

// SetEtcVersion is EtcVersion Setter
// ms级时间戳
func (r *AlibabaalihouseexistinghomehouseupselfAPIRequest) SetEtcVersion(_etcVersion int64) error {
	r._etcVersion = _etcVersion
	r.Set("etc_version", _etcVersion)
	return nil
}

// GetEtcVersion EtcVersion Getter
func (r AlibabaalihouseexistinghomehouseupselfAPIRequest) GetEtcVersion() int64 {
	return r._etcVersion
}
