package flight

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPolicyDomfareFlowdataAPIRequest) Reset() {
	r._compareFlowDataQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyDomfareFlowdataAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.domfare.flowdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyDomfareFlowdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicyDomfareFlowdataAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripPolicyDomfareFlowdataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPolicyDomfareFlowdataRequest()
	},
}

// GetAlitripPolicyDomfareFlowdataRequest 从 sync.Pool 获取 AlitripPolicyDomfareFlowdataAPIRequest
func GetAlitripPolicyDomfareFlowdataAPIRequest() *AlitripPolicyDomfareFlowdataAPIRequest {
	return poolAlitripPolicyDomfareFlowdataAPIRequest.Get().(*AlitripPolicyDomfareFlowdataAPIRequest)
}

// ReleaseAlitripPolicyDomfareFlowdataAPIRequest 将 AlitripPolicyDomfareFlowdataAPIRequest 放入 sync.Pool
func ReleaseAlitripPolicyDomfareFlowdataAPIRequest(v *AlitripPolicyDomfareFlowdataAPIRequest) {
	v.Reset()
	poolAlitripPolicyDomfareFlowdataAPIRequest.Put(v)
}
