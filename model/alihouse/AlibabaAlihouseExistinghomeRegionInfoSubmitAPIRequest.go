package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest 商圈专家信息同步 API请求
// alibaba.alihouse.existinghome.region.info.submit
//
// 商圈专家信息同步
type AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest struct {
	model.Params
	// 商圈信息
	_regionExpertDto *RegionExpertDto
}

// NewAlibabaAlihouseExistinghomeRegionInfoSubmitRequest 初始化AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest对象
func NewAlibabaAlihouseExistinghomeRegionInfoSubmitRequest() *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest {
	return &AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest) Reset() {
	r._regionExpertDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.region.info.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegionExpertDto is RegionExpertDto Setter
// 商圈信息
func (r *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest) SetRegionExpertDto(_regionExpertDto *RegionExpertDto) error {
	r._regionExpertDto = _regionExpertDto
	r.Set("region_expert_dto", _regionExpertDto)
	return nil
}

// GetRegionExpertDto RegionExpertDto Getter
func (r AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest) GetRegionExpertDto() *RegionExpertDto {
	return r._regionExpertDto
}

var poolAlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeRegionInfoSubmitRequest()
	},
}

// GetAlibabaAlihouseExistinghomeRegionInfoSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest
func GetAlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest() *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest {
	return poolAlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest.Get().(*AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest 将 AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest(v *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeRegionInfoSubmitAPIRequest.Put(v)
}
