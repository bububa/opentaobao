package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeregioninfosubmitAPIRequest 商圈专家信息同步 API请求
// alibaba.alihouse.existinghome.region.info.submit
//
// 商圈专家信息同步
type AlibabaalihouseexistinghomeregioninfosubmitAPIRequest struct {
	model.Params
	// 商圈信息
	_regionExpertDto *RegionExpertDto
}

// NewAlibabaalihouseexistinghomeregioninfosubmitRequest 初始化AlibabaalihouseexistinghomeregioninfosubmitAPIRequest对象
func NewAlibabaalihouseexistinghomeregioninfosubmitRequest() *AlibabaalihouseexistinghomeregioninfosubmitAPIRequest {
	return &AlibabaalihouseexistinghomeregioninfosubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomeregioninfosubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.region.info.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomeregioninfosubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomeregioninfosubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRegionExpertDto is RegionExpertDto Setter
// 商圈信息
func (r *AlibabaalihouseexistinghomeregioninfosubmitAPIRequest) SetRegionExpertDto(_regionExpertDto *RegionExpertDto) error {
	r._regionExpertDto = _regionExpertDto
	r.Set("region_expert_dto", _regionExpertDto)
	return nil
}

// GetRegionExpertDto RegionExpertDto Getter
func (r AlibabaalihouseexistinghomeregioninfosubmitAPIRequest) GetRegionExpertDto() *RegionExpertDto {
	return r._regionExpertDto
}
