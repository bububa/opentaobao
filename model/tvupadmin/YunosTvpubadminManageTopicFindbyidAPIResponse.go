package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopicfindbyidAPIResponse 根据id获取专题信息 API返回值
// yunos.tvpubadmin.manage.topic.findbyid
//
// 根据id获取专题信息
type YunostvpubadminmanagetopicfindbyidAPIResponse struct {
	model.CommonResponse
	YunostvpubadminmanagetopicfindbyidAPIResponseModel
}

// YunostvpubadminmanagetopicfindbyidAPIResponseModel is 根据id获取专题信息 成功返回结果
type YunostvpubadminmanagetopicfindbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_findbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
