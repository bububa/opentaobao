package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest 提交小区动态信息 API请求
// alibaba.alihouse.newhome.dynamicdata.submit
//
// 提交小区动态信息
type AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest struct {
	model.Params
	// 小区价格变动实体
	_dynamicDataDto *DynamicDataDto
}

// NewAlibabaAlihouseNewhomeDynamicdataSubmitRequest 初始化AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeDynamicdataSubmitRequest() *AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest) Reset() {
	r._dynamicDataDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.dynamicdata.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDynamicDataDto is DynamicDataDto Setter
// 小区价格变动实体
func (r *AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest) SetDynamicDataDto(_dynamicDataDto *DynamicDataDto) error {
	r._dynamicDataDto = _dynamicDataDto
	r.Set("dynamic_data_dto", _dynamicDataDto)
	return nil
}

// GetDynamicDataDto DynamicDataDto Getter
func (r AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest) GetDynamicDataDto() *DynamicDataDto {
	return r._dynamicDataDto
}

var poolAlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeDynamicdataSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeDynamicdataSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest
func GetAlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest() *AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest 将 AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest(v *AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest.Put(v)
}
