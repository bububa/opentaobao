package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiclistAPIResponse 专题内容操作列表 API返回值
// yunos.tvpubadmin.manage.topic.list
//
// 获取外部可操作编辑的专题列表
type YunostvpubadminmanagetopiclistAPIResponse struct {
	model.CommonResponse
	YunostvpubadminmanagetopiclistAPIResponseModel
}

// YunostvpubadminmanagetopiclistAPIResponseModel is 专题内容操作列表 成功返回结果
type YunostvpubadminmanagetopiclistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
