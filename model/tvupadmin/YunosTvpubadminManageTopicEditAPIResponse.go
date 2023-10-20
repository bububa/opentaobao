package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiceditAPIResponse 编辑专题 API返回值
// yunos.tvpubadmin.manage.topic.edit
//
// 编辑专题
type YunostvpubadminmanagetopiceditAPIResponse struct {
	model.CommonResponse
	YunostvpubadminmanagetopiceditAPIResponseModel
}

// YunostvpubadminmanagetopiceditAPIResponseModel is 编辑专题 成功返回结果
type YunostvpubadminmanagetopiceditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunostvpubadminmanagetopiceditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
