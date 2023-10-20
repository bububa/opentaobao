package normalvisa

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelvisasignsendAPIResponse 签证批量申请人送签接口 API返回值
// alitrip.travel.visa.sign.send
//
// 签证批量申请人送签接口，用于商家批量送签。
type AlitriptravelvisasignsendAPIResponse struct {
	model.CommonResponse
	AlitriptravelvisasignsendAPIResponseModel
}

// AlitriptravelvisasignsendAPIResponseModel is 签证批量申请人送签接口 成功返回结果
type AlitriptravelvisasignsendAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_visa_sign_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 批次信息
	BatchInfos []BatchInfo `json:"batch_infos,omitempty" xml:"batch_infos>batch_info,omitempty"`
	// 失败信息
	FailInfos []SendSignFailInfo `json:"fail_infos,omitempty" xml:"fail_infos>send_sign_fail_info,omitempty"`
}
