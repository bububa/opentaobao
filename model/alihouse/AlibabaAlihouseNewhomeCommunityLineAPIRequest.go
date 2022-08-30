package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCommunityLineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.community.line"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCommunityLineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
