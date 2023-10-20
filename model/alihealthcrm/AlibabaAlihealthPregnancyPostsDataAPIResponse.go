package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyPostsDataAPIResponse 发回帖子信息同步 API返回值
// alibaba.alihealth.pregnancy.posts.data
//
// 发回帖子信息同步
type AlibabaAlihealthPregnancyPostsDataAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPregnancyPostsDataAPIResponseModel
}

// AlibabaAlihealthPregnancyPostsDataAPIResponseModel is 发回帖子信息同步 成功返回结果
type AlibabaAlihealthPregnancyPostsDataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_posts_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
