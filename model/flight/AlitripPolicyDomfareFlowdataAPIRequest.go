package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyDomfareFlowdataAPIRequest 比价工具流量详情 API请求
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
type AlitripPolicyDomfareFlowdataAPIRequest struct {
	model.Params
	// 入参
	_compareFlowDataQueryDTO *CompareFlowDataQueryDto
}

// NewAlitripPolicyDomfareFlowdataRequest 初始化AlitripPolicyDomfareFlowdataAPIRequest对象
func NewAlitripPolicyDomfareFlowdataRequest() *AlitripPolicyDomfareFlowdataAPIRequest {
	return &AlitripPolicyDomfareFlowdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyDomfareFlowdataAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.domfare.flowdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyDomfareFlowdataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCompareFlowDataQueryDTO is CompareFlowDataQueryDTO Setter
// 入参
func (r *AlitripPolicyDomfareFlowdataAPIRequest) SetCompareFlowDataQueryDTO(_compareFlowDataQueryDTO *CompareFlowDataQueryDto) error {
	r._compareFlowDataQueryDTO = _compareFlowDataQueryDTO
	r.Set("compare_flow_data_query_d_t_o", _compareFlowDataQueryDTO)
	return nil
}

// GetCompareFlowDataQueryDTO CompareFlowDataQueryDTO Getter
func (r AlitripPolicyDomfareFlowdataAPIRequest) GetCompareFlowDataQueryDTO() *CompareFlowDataQueryDto {
	return r._compareFlowDataQueryDTO
}
