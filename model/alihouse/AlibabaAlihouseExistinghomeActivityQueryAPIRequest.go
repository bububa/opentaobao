package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeActivityQueryAPIRequest 五福活动经纪人获取 API请求
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
type AlibabaAlihouseExistinghomeActivityQueryAPIRequest struct {
	model.Params
	// 查询的参数
	_queryInfo *ActivityQueryInfoDto
}

// NewAlibabaAlihouseExistinghomeActivityQueryRequest 初始化AlibabaAlihouseExistinghomeActivityQueryAPIRequest对象
func NewAlibabaAlihouseExistinghomeActivityQueryRequest() *AlibabaAlihouseExistinghomeActivityQueryAPIRequest {
	return &AlibabaAlihouseExistinghomeActivityQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeActivityQueryAPIRequest) Reset() {
	r._queryInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryInfo is QueryInfo Setter
// 查询的参数
func (r *AlibabaAlihouseExistinghomeActivityQueryAPIRequest) SetQueryInfo(_queryInfo *ActivityQueryInfoDto) error {
	r._queryInfo = _queryInfo
	r.Set("query_info", _queryInfo)
	return nil
}

// GetQueryInfo QueryInfo Getter
func (r AlibabaAlihouseExistinghomeActivityQueryAPIRequest) GetQueryInfo() *ActivityQueryInfoDto {
	return r._queryInfo
}

var poolAlibabaAlihouseExistinghomeActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeActivityQueryRequest()
	},
}

// GetAlibabaAlihouseExistinghomeActivityQueryRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeActivityQueryAPIRequest
func GetAlibabaAlihouseExistinghomeActivityQueryAPIRequest() *AlibabaAlihouseExistinghomeActivityQueryAPIRequest {
	return poolAlibabaAlihouseExistinghomeActivityQueryAPIRequest.Get().(*AlibabaAlihouseExistinghomeActivityQueryAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeActivityQueryAPIRequest 将 AlibabaAlihouseExistinghomeActivityQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeActivityQueryAPIRequest(v *AlibabaAlihouseExistinghomeActivityQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeActivityQueryAPIRequest.Put(v)
}
