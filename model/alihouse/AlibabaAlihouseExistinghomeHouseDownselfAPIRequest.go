package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeHouseDownselfAPIRequest 房源信息下架 API请求
// alibaba.alihouse.existinghome.house.downself
//
// 房源信息下架
type AlibabaAlihouseExistinghomeHouseDownselfAPIRequest struct {
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

// NewAlibabaAlihouseExistinghomeHouseDownselfRequest 初始化AlibabaAlihouseExistinghomeHouseDownselfAPIRequest对象
func NewAlibabaAlihouseExistinghomeHouseDownselfRequest() *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest {
	return &AlibabaAlihouseExistinghomeHouseDownselfAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) Reset() {
	r._communityOuterId = ""
	r._outerId = ""
	r._outerEntrustId = ""
	r._businessType = 0
	r._isAsync = 0
	r._etcVersion = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.downself"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommunityOuterId is CommunityOuterId Setter
// 外部小区id
func (r *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) SetCommunityOuterId(_communityOuterId string) error {
	r._communityOuterId = _communityOuterId
	r.Set("community_outer_id", _communityOuterId)
	return nil
}

// GetCommunityOuterId CommunityOuterId Getter
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetCommunityOuterId() string {
	return r._communityOuterId
}

// SetOuterId is OuterId Setter
// 外部房源id
func (r *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterEntrustId is OuterEntrustId Setter
// 外部委托id
func (r *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) SetOuterEntrustId(_outerEntrustId string) error {
	r._outerEntrustId = _outerEntrustId
	r.Set("outer_entrust_id", _outerEntrustId)
	return nil
}

// GetOuterEntrustId OuterEntrustId Getter
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetOuterEntrustId() string {
	return r._outerEntrustId
}

// SetBusinessType is BusinessType Setter
// 1-新房，2-小区，3-公寓。与楼盘的枚举对齐，目的是查询小区或者是公寓数据
func (r *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetBusinessType() int64 {
	return r._businessType
}

// SetIsAsync is IsAsync Setter
// 是否异步，0同步，1异步
func (r *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) SetIsAsync(_isAsync int64) error {
	r._isAsync = _isAsync
	r.Set("is_async", _isAsync)
	return nil
}

// GetIsAsync IsAsync Getter
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetIsAsync() int64 {
	return r._isAsync
}

// SetEtcVersion is EtcVersion Setter
// ms级时间戳
func (r *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) SetEtcVersion(_etcVersion int64) error {
	r._etcVersion = _etcVersion
	r.Set("etc_version", _etcVersion)
	return nil
}

// GetEtcVersion EtcVersion Getter
func (r AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) GetEtcVersion() int64 {
	return r._etcVersion
}

var poolAlibabaAlihouseExistinghomeHouseDownselfAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeHouseDownselfRequest()
	},
}

// GetAlibabaAlihouseExistinghomeHouseDownselfRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeHouseDownselfAPIRequest
func GetAlibabaAlihouseExistinghomeHouseDownselfAPIRequest() *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest {
	return poolAlibabaAlihouseExistinghomeHouseDownselfAPIRequest.Get().(*AlibabaAlihouseExistinghomeHouseDownselfAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeHouseDownselfAPIRequest 将 AlibabaAlihouseExistinghomeHouseDownselfAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeHouseDownselfAPIRequest(v *AlibabaAlihouseExistinghomeHouseDownselfAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeHouseDownselfAPIRequest.Put(v)
}
