package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChildLeafnodeGetAPIResponse
获取少儿大厅二级类目列表 API返回值
yunos.tvpubadmin.content.child.leafnode.get

获取少儿大厅二级类目列表的接口 */
type YunosTvpubadminContentChildLeafnodeGetAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildLeafnodeGetAPIResponseModel
}

// YunosTvpubadminContentChildLeafnodeGetAPIResponseModel is 获取少儿大厅二级类目列表 成功返回结果
type YunosTvpubadminContentChildLeafnodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_leafnode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目列表
	ChildNodeVoList []ChildNodeVo `json:"child_node_vo_list,omitempty" xml:"child_node_vo_list>child_node_vo,omitempty"`
}
