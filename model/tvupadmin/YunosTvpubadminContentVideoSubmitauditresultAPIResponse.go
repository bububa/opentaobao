package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentvideosubmitauditresultAPIResponse 迎客松提交视频审核结果 API返回值
// yunos.tvpubadmin.content.video.submitauditresult
//
// 迎客松提交视频审核结果
type YunostvpubadmincontentvideosubmitauditresultAPIResponse struct {
	model.CommonResponse
	YunostvpubadmincontentvideosubmitauditresultAPIResponseModel
}

// YunostvpubadmincontentvideosubmitauditresultAPIResponseModel is 迎客松提交视频审核结果 成功返回结果
type YunostvpubadmincontentvideosubmitauditresultAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_video_submitauditresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}
