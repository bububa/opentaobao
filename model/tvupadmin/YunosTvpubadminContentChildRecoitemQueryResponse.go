package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询少儿大厅推荐内容列表 APIResponse
yunos.tvpubadmin.content.child.recoitem.query

查询少儿大厅推荐内容列表
*/
type YunosTvpubadminContentChildRecoitemQueryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentChildRecoitemQueryResponse
}

type YunosTvpubadminContentChildRecoitemQueryResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_recoitem_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 搜索结果
    
    Result   *PageVo `json:"result,omitempty" xml:"result,omitempty"`

    
}
