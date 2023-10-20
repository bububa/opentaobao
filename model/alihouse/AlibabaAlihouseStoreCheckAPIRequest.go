package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseStoreCheckAPIRequest 门店对账查询工具 API请求
// alibaba.alihouse.store.check
//
// 门店对账查询工具
type AlibabaAlihouseStoreCheckAPIRequest struct {
	model.Params
	// 外部id列表
	_outerIds []string
}

// NewAlibabaAlihouseStoreCheckRequest 初始化AlibabaAlihouseStoreCheckAPIRequest对象
func NewAlibabaAlihouseStoreCheckRequest() *AlibabaAlihouseStoreCheckAPIRequest {
	return &AlibabaAlihouseStoreCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseStoreCheckAPIRequest) Reset() {
	r._outerIds = r._outerIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseStoreCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.store.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseStoreCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseStoreCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterIds is OuterIds Setter
// 外部id列表
func (r *AlibabaAlihouseStoreCheckAPIRequest) SetOuterIds(_outerIds []string) error {
	r._outerIds = _outerIds
	r.Set("outer_ids", _outerIds)
	return nil
}

// GetOuterIds OuterIds Getter
func (r AlibabaAlihouseStoreCheckAPIRequest) GetOuterIds() []string {
	return r._outerIds
}

var poolAlibabaAlihouseStoreCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseStoreCheckRequest()
	},
}

// GetAlibabaAlihouseStoreCheckRequest 从 sync.Pool 获取 AlibabaAlihouseStoreCheckAPIRequest
func GetAlibabaAlihouseStoreCheckAPIRequest() *AlibabaAlihouseStoreCheckAPIRequest {
	return poolAlibabaAlihouseStoreCheckAPIRequest.Get().(*AlibabaAlihouseStoreCheckAPIRequest)
}

// ReleaseAlibabaAlihouseStoreCheckAPIRequest 将 AlibabaAlihouseStoreCheckAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseStoreCheckAPIRequest(v *AlibabaAlihouseStoreCheckAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseStoreCheckAPIRequest.Put(v)
}
