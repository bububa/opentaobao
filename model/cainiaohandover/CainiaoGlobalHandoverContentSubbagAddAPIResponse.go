package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandovercontentsubbagaddAPIResponse 预约单下追加大包 API返回值
// cainiao.global.handover.content.subbag.add
//
// 预约单下追加大包
type CainiaoglobalhandovercontentsubbagaddAPIResponse struct {
	model.CommonResponse
	CainiaoglobalhandovercontentsubbagaddAPIResponseModel
}

// CainiaoglobalhandovercontentsubbagaddAPIResponseModel is 预约单下追加大包 成功返回结果
type CainiaoglobalhandovercontentsubbagaddAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_content_subbag_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求响应
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}
