package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandovercontentqueryAPIResponse 查询大包详情 API返回值
// cainiao.global.handover.content.query
//
// 查询大包详情
type CainiaoglobalhandovercontentqueryAPIResponse struct {
	model.CommonResponse
	CainiaoglobalhandovercontentqueryAPIResponseModel
}

// CainiaoglobalhandovercontentqueryAPIResponseModel is 查询大包详情 成功返回结果
type CainiaoglobalhandovercontentqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_content_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求响应
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}
