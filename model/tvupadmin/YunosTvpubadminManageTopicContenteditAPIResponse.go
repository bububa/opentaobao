package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiccontenteditAPIResponse 专题关联内容编辑 API返回值
// yunos.tvpubadmin.manage.topic.contentedit
//
// 编辑专题关联的内容
type YunostvpubadminmanagetopiccontenteditAPIResponse struct {
	model.CommonResponse
	YunostvpubadminmanagetopiccontenteditAPIResponseModel
}

// YunostvpubadminmanagetopiccontenteditAPIResponseModel is 专题关联内容编辑 成功返回结果
type YunostvpubadminmanagetopiccontenteditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_contentedit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作返回结果
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
}
