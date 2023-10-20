package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiccontentdeleteAPIResponse 删除专题下内容 API返回值
// yunos.tvpubadmin.manage.topic.contentdelete
//
// 删除专题下内容信息
type YunostvpubadminmanagetopiccontentdeleteAPIResponse struct {
	model.CommonResponse
	YunostvpubadminmanagetopiccontentdeleteAPIResponseModel
}

// YunostvpubadminmanagetopiccontentdeleteAPIResponseModel is 删除专题下内容 成功返回结果
type YunostvpubadminmanagetopiccontentdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentdelete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *YunostvpubadminmanagetopiccontentdeleteTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
