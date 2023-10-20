package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandovercloudprintgetAPIResponse 获取面单云打印数据 API返回值
// cainiao.global.handover.cloudprint.get
//
// 提供给ISV通过该接口获取面单云打印数据
type CainiaoglobalhandovercloudprintgetAPIResponse struct {
	model.CommonResponse
	CainiaoglobalhandovercloudprintgetAPIResponseModel
}

// CainiaoglobalhandovercloudprintgetAPIResponseModel is 获取面单云打印数据 成功返回结果
type CainiaoglobalhandovercloudprintgetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_cloudprint_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}
