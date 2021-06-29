package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取少儿大厅二级类目列表 APIResponse
yunos.tvpubadmin.content.child.leafnode.get

获取少儿大厅二级类目列表的接口
*/
type YunosTvpubadminContentChildLeafnodeGetAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentChildLeafnodeGetResponse
}

type YunosTvpubadminContentChildLeafnodeGetResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_leafnode_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 类目列表
    
    ChildNodeVoList   []ChildNodeVo `json:"child_node_vo_list,omitempty" xml:"child_node_vo_list>child_node_vo,omitempty"`
    
    
}
