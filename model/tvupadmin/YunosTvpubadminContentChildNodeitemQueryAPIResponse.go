package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildNodeitemQueryAPIResponse 查询少儿大厅类目内容 API返回值
// yunos.tvpubadmin.content.child.nodeitem.query
//
// 查询少儿大厅类目内容信息
type YunosTvpubadminContentChildNodeitemQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildNodeitemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildNodeitemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentChildNodeitemQueryAPIResponseModel).Reset()
}

// YunosTvpubadminContentChildNodeitemQueryAPIResponseModel is 查询少儿大厅类目内容 成功返回结果
type YunosTvpubadminContentChildNodeitemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_nodeitem_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	PageVo *PageVo `json:"page_vo,omitempty" xml:"page_vo,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildNodeitemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PageVo = nil
}

var poolYunosTvpubadminContentChildNodeitemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentChildNodeitemQueryAPIResponse)
	},
}

// GetYunosTvpubadminContentChildNodeitemQueryAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentChildNodeitemQueryAPIResponse
func GetYunosTvpubadminContentChildNodeitemQueryAPIResponse() *YunosTvpubadminContentChildNodeitemQueryAPIResponse {
	return poolYunosTvpubadminContentChildNodeitemQueryAPIResponse.Get().(*YunosTvpubadminContentChildNodeitemQueryAPIResponse)
}

// ReleaseYunosTvpubadminContentChildNodeitemQueryAPIResponse 将 YunosTvpubadminContentChildNodeitemQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentChildNodeitemQueryAPIResponse(v *YunosTvpubadminContentChildNodeitemQueryAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentChildNodeitemQueryAPIResponse.Put(v)
}
