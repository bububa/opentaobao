package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiEcSearchProjectSearchAPIResponse 大麦电商对外搜索服务 API返回值
// alibaba.damai.ec.search.project.search
//
// 大麦电商对外搜索服务
type AlibabaDamaiEcSearchProjectSearchAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiEcSearchProjectSearchAPIResponseModel
}

// AlibabaDamaiEcSearchProjectSearchAPIResponseModel is 大麦电商对外搜索服务 成功返回结果
type AlibabaDamaiEcSearchProjectSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_ec_search_project_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MpResult `json:"result,omitempty" xml:"result,omitempty"`
}
