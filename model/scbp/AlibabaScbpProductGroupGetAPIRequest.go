package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductGroupGetAPIRequest 查询指定产品分组的下一层子分组 API请求
// alibaba.scbp.product.group.get
//
// 查询指定产品分组的下一层子分组
type AlibabaScbpProductGroupGetAPIRequest struct {
	model.Params
	// 产品分组标识，null表示查询第一层分组
	_groupId string
}

// NewAlibabaScbpProductGroupGetRequest 初始化AlibabaScbpProductGroupGetAPIRequest对象
func NewAlibabaScbpProductGroupGetRequest() *AlibabaScbpProductGroupGetAPIRequest {
	return &AlibabaScbpProductGroupGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpProductGroupGetAPIRequest) Reset() {
	r._groupId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductGroupGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductGroupGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpProductGroupGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 产品分组标识，null表示查询第一层分组
func (r *AlibabaScbpProductGroupGetAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaScbpProductGroupGetAPIRequest) GetGroupId() string {
	return r._groupId
}

var poolAlibabaScbpProductGroupGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpProductGroupGetRequest()
	},
}

// GetAlibabaScbpProductGroupGetRequest 从 sync.Pool 获取 AlibabaScbpProductGroupGetAPIRequest
func GetAlibabaScbpProductGroupGetAPIRequest() *AlibabaScbpProductGroupGetAPIRequest {
	return poolAlibabaScbpProductGroupGetAPIRequest.Get().(*AlibabaScbpProductGroupGetAPIRequest)
}

// ReleaseAlibabaScbpProductGroupGetAPIRequest 将 AlibabaScbpProductGroupGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpProductGroupGetAPIRequest(v *AlibabaScbpProductGroupGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpProductGroupGetAPIRequest.Put(v)
}
