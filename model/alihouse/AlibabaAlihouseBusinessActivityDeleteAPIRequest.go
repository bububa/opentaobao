package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseBusinessActivityDeleteAPIRequest 电商券活动删除 API请求
// alibaba.alihouse.business.activity.delete
//
// 电商券活动删除
type AlibabaAlihouseBusinessActivityDeleteAPIRequest struct {
	model.Params
	// 数据
	_baiYiActivityDataDto *BusinessActivityDataDto
}

// NewAlibabaAlihouseBusinessActivityDeleteRequest 初始化AlibabaAlihouseBusinessActivityDeleteAPIRequest对象
func NewAlibabaAlihouseBusinessActivityDeleteRequest() *AlibabaAlihouseBusinessActivityDeleteAPIRequest {
	return &AlibabaAlihouseBusinessActivityDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseBusinessActivityDeleteAPIRequest) Reset() {
	r._baiYiActivityDataDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseBusinessActivityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.business.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseBusinessActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseBusinessActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaiYiActivityDataDto is BaiYiActivityDataDto Setter
// 数据
func (r *AlibabaAlihouseBusinessActivityDeleteAPIRequest) SetBaiYiActivityDataDto(_baiYiActivityDataDto *BusinessActivityDataDto) error {
	r._baiYiActivityDataDto = _baiYiActivityDataDto
	r.Set("bai_yi_activity_data_dto", _baiYiActivityDataDto)
	return nil
}

// GetBaiYiActivityDataDto BaiYiActivityDataDto Getter
func (r AlibabaAlihouseBusinessActivityDeleteAPIRequest) GetBaiYiActivityDataDto() *BusinessActivityDataDto {
	return r._baiYiActivityDataDto
}

var poolAlibabaAlihouseBusinessActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseBusinessActivityDeleteRequest()
	},
}

// GetAlibabaAlihouseBusinessActivityDeleteRequest 从 sync.Pool 获取 AlibabaAlihouseBusinessActivityDeleteAPIRequest
func GetAlibabaAlihouseBusinessActivityDeleteAPIRequest() *AlibabaAlihouseBusinessActivityDeleteAPIRequest {
	return poolAlibabaAlihouseBusinessActivityDeleteAPIRequest.Get().(*AlibabaAlihouseBusinessActivityDeleteAPIRequest)
}

// ReleaseAlibabaAlihouseBusinessActivityDeleteAPIRequest 将 AlibabaAlihouseBusinessActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseBusinessActivityDeleteAPIRequest(v *AlibabaAlihouseBusinessActivityDeleteAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseBusinessActivityDeleteAPIRequest.Put(v)
}
