package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentshowsetexemptauditAPIResponse 迎客松节目设置免审开关 API返回值
// yunos.tvpubadmin.content.show.setexemptaudit
//
// 迎客松节目设置免审开关
type YunostvpubadmincontentshowsetexemptauditAPIResponse struct {
	model.CommonResponse
	YunostvpubadmincontentshowsetexemptauditAPIResponseModel
}

// YunostvpubadmincontentshowsetexemptauditAPIResponseModel is 迎客松节目设置免审开关 成功返回结果
type YunostvpubadmincontentshowsetexemptauditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_setexemptaudit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设置免审
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}
