package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomeregioninfosubmitAPIResponse 商圈专家信息同步 API返回值
// alibaba.alihouse.existinghome.region.info.submit
//
// 商圈专家信息同步
type AlibabaalihouseexistinghomeregioninfosubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomeregioninfosubmitAPIResponseModel
}

// AlibabaalihouseexistinghomeregioninfosubmitAPIResponseModel is 商圈专家信息同步 成功返回结果
type AlibabaalihouseexistinghomeregioninfosubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_region_info_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihouseexistinghomeregioninfosubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
