package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentChildRecoitemQueryAPIResponse
查询少儿大厅推荐内容列表 API返回值
yunos.tvpubadmin.content.child.recoitem.query

查询少儿大厅推荐内容列表 */
type YunosTvpubadminContentChildRecoitemQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildRecoitemQueryAPIResponseModel
}

// YunosTvpubadminContentChildRecoitemQueryAPIResponseModel is 查询少儿大厅推荐内容列表 成功返回结果
type YunosTvpubadminContentChildRecoitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_recoitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索结果
	Result *PageVo `json:"result,omitempty" xml:"result,omitempty"`
}
