package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicydomfareflowdataAPIRequest 比价工具流量详情 API请求
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
type AlitrippolicydomfareflowdataAPIRequest struct {
	model.Params
	// 入参
	_compareFlowDataQueryDTO *CompareFlowDataQueryDto
}

// NewAlitrippolicydomfareflowdataRequest 初始化AlitrippolicydomfareflowdataAPIRequest对象
func NewAlitrippolicydomfareflowdataRequest() *AlitrippolicydomfareflowdataAPIRequest {
	return &AlitrippolicydomfareflowdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitrippolicydomfareflowdataAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.domfare.flowdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitrippolicydomfareflowdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitrippolicydomfareflowdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompareFlowDataQueryDTO is CompareFlowDataQueryDTO Setter
// 入参
func (r *AlitrippolicydomfareflowdataAPIRequest) SetCompareFlowDataQueryDTO(_compareFlowDataQueryDTO *CompareFlowDataQueryDto) error {
	r._compareFlowDataQueryDTO = _compareFlowDataQueryDTO
	r.Set("compare_flow_data_query_d_t_o", _compareFlowDataQueryDTO)
	return nil
}

// GetCompareFlowDataQueryDTO CompareFlowDataQueryDTO Getter
func (r AlitrippolicydomfareflowdataAPIRequest) GetCompareFlowDataQueryDTO() *CompareFlowDataQueryDto {
	return r._compareFlowDataQueryDTO
}
