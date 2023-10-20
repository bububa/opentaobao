package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductGroupGetAPIRequest 分组信息获取 API请求
// alibaba.icbu.product.group.get
//
// 分组信息获取
type AlibabaIcbuProductGroupGetAPIRequest struct {
	model.Params
	// 补充信息
	_extraContext string
	// 分组ID，传-1获得所有一级分组
	_groupId int64
}

// NewAlibabaIcbuProductGroupGetRequest 初始化AlibabaIcbuProductGroupGetAPIRequest对象
func NewAlibabaIcbuProductGroupGetRequest() *AlibabaIcbuProductGroupGetAPIRequest {
	return &AlibabaIcbuProductGroupGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductGroupGetAPIRequest) Reset() {
	r._extraContext = ""
	r._groupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGroupGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGroupGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductGroupGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraContext is ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuProductGroupGetAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaIcbuProductGroupGetAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetGroupId is GroupId Setter
// 分组ID，传-1获得所有一级分组
func (r *AlibabaIcbuProductGroupGetAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaIcbuProductGroupGetAPIRequest) GetGroupId() int64 {
	return r._groupId
}

var poolAlibabaIcbuProductGroupGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductGroupGetRequest()
	},
}

// GetAlibabaIcbuProductGroupGetRequest 从 sync.Pool 获取 AlibabaIcbuProductGroupGetAPIRequest
func GetAlibabaIcbuProductGroupGetAPIRequest() *AlibabaIcbuProductGroupGetAPIRequest {
	return poolAlibabaIcbuProductGroupGetAPIRequest.Get().(*AlibabaIcbuProductGroupGetAPIRequest)
}

// ReleaseAlibabaIcbuProductGroupGetAPIRequest 将 AlibabaIcbuProductGroupGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductGroupGetAPIRequest(v *AlibabaIcbuProductGroupGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductGroupGetAPIRequest.Put(v)
}
