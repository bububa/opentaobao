package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildRootnodeGetAPIResponse 获取少儿大厅根类目接口 API返回值
// yunos.tvpubadmin.content.child.rootnode.get
//
// 通过此接口可获取少儿大厅根类目列表
type YunosTvpubadminContentChildRootnodeGetAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildRootnodeGetAPIResponseModel
}

// YunosTvpubadminContentChildRootnodeGetAPIResponseModel is 获取少儿大厅根类目接口 成功返回结果
type YunosTvpubadminContentChildRootnodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_rootnode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目列表
	ChildNodeVoList []ChildNodeVo `json:"child_node_vo_list,omitempty" xml:"child_node_vo_list>child_node_vo,omitempty"`
}
