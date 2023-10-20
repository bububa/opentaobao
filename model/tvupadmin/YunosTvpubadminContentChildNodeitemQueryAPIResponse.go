package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchildnodeitemqueryAPIResponse 查询少儿大厅类目内容 API返回值
// yunos.tvpubadmin.content.child.nodeitem.query
//
// 查询少儿大厅类目内容信息
type YunostvpubadmincontentchildnodeitemqueryAPIResponse struct {
	model.CommonResponse
	YunostvpubadmincontentchildnodeitemqueryAPIResponseModel
}

// YunostvpubadmincontentchildnodeitemqueryAPIResponseModel is 查询少儿大厅类目内容 成功返回结果
type YunostvpubadmincontentchildnodeitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_nodeitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	PageVo *PageVo `json:"page_vo,omitempty" xml:"page_vo,omitempty"`
}
