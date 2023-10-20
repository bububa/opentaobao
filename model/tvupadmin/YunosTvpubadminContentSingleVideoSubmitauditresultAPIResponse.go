package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponse 单视频审核提交审核结果 API返回值
// yunos.tvpubadmin.content.single.video.submitauditresult
//
// 单视频审核提交审核结果
type YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponseModel
}

// YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponseModel is 单视频审核提交审核结果 成功返回结果
type YunosTvpubadminContentSingleVideoSubmitauditresultAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_single_video_submitauditresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}
