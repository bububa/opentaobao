package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCommunityLineAPIRequest 小区上下架 API请求
// alibaba.alihouse.newhome.community.line
//
// 小区上下架
type AlibabaAlihouseNewhomeCommunityLineAPIRequest struct {
	model.Params
	// 外部小区ID
	_outerId string
	// 0-下架 1-上架
	_type *model.File
}

// NewAlibabaAlihouseNewhomeCommunityLineRequest 初始化AlibabaAlihouseNewhomeCommunityLineAPIRequest对象
func NewAlibabaAlihouseNewhomeCommunityLineRequest() *AlibabaAlihouseNewhomeCommunityLineAPIRequest {
	return &AlibabaAlihouseNewhomeCommunityLineAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeCommunityLineAPIRequest) Reset() {
	r._outerId = ""
	r._type = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCommunityLineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.community.line"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCommunityLineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeCommunityLineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部小区ID
func (r *AlibabaAlihouseNewhomeCommunityLineAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeCommunityLineAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetType is Type Setter
// 0-下架 1-上架
func (r *AlibabaAlihouseNewhomeCommunityLineAPIRequest) SetType(_type *model.File) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihouseNewhomeCommunityLineAPIRequest) GetType() *model.File {
	return r._type
}

var poolAlibabaAlihouseNewhomeCommunityLineAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeCommunityLineRequest()
	},
}

// GetAlibabaAlihouseNewhomeCommunityLineRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeCommunityLineAPIRequest
func GetAlibabaAlihouseNewhomeCommunityLineAPIRequest() *AlibabaAlihouseNewhomeCommunityLineAPIRequest {
	return poolAlibabaAlihouseNewhomeCommunityLineAPIRequest.Get().(*AlibabaAlihouseNewhomeCommunityLineAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeCommunityLineAPIRequest 将 AlibabaAlihouseNewhomeCommunityLineAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCommunityLineAPIRequest(v *AlibabaAlihouseNewhomeCommunityLineAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCommunityLineAPIRequest.Put(v)
}
