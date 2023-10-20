package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
