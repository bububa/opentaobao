package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildRecoitemQueryAPIResponse 查询少儿大厅推荐内容列表 API返回值
// yunos.tvpubadmin.content.child.recoitem.query
//
// 查询少儿大厅推荐内容列表
type YunosTvpubadminContentChildRecoitemQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildRecoitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildRecoitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentChildRecoitemQueryAPIResponseModel).Reset()
}

// YunosTvpubadminContentChildRecoitemQueryAPIResponseModel is 查询少儿大厅推荐内容列表 成功返回结果
type YunosTvpubadminContentChildRecoitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_recoitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索结果
	Result *PageVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildRecoitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTvpubadminContentChildRecoitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentChildRecoitemQueryAPIResponse)
	},
}

// GetYunosTvpubadminContentChildRecoitemQueryAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentChildRecoitemQueryAPIResponse
func GetYunosTvpubadminContentChildRecoitemQueryAPIResponse() *YunosTvpubadminContentChildRecoitemQueryAPIResponse {
	return poolYunosTvpubadminContentChildRecoitemQueryAPIResponse.Get().(*YunosTvpubadminContentChildRecoitemQueryAPIResponse)
}

// ReleaseYunosTvpubadminContentChildRecoitemQueryAPIResponse 将 YunosTvpubadminContentChildRecoitemQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentChildRecoitemQueryAPIResponse(v *YunosTvpubadminContentChildRecoitemQueryAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentChildRecoitemQueryAPIResponse.Put(v)
}
