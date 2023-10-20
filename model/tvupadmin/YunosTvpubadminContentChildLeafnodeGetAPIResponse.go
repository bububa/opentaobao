package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildLeafnodeGetAPIResponse 获取少儿大厅二级类目列表 API返回值
// yunos.tvpubadmin.content.child.leafnode.get
//
// 获取少儿大厅二级类目列表的接口
type YunosTvpubadminContentChildLeafnodeGetAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentChildLeafnodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildLeafnodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentChildLeafnodeGetAPIResponseModel).Reset()
}

// YunosTvpubadminContentChildLeafnodeGetAPIResponseModel is 获取少儿大厅二级类目列表 成功返回结果
type YunosTvpubadminContentChildLeafnodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_leafnode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目列表
	ChildNodeVoList []ChildNodeVo `json:"child_node_vo_list,omitempty" xml:"child_node_vo_list>child_node_vo,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentChildLeafnodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ChildNodeVoList = m.ChildNodeVoList[:0]
}

var poolYunosTvpubadminContentChildLeafnodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentChildLeafnodeGetAPIResponse)
	},
}

// GetYunosTvpubadminContentChildLeafnodeGetAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentChildLeafnodeGetAPIResponse
func GetYunosTvpubadminContentChildLeafnodeGetAPIResponse() *YunosTvpubadminContentChildLeafnodeGetAPIResponse {
	return poolYunosTvpubadminContentChildLeafnodeGetAPIResponse.Get().(*YunosTvpubadminContentChildLeafnodeGetAPIResponse)
}

// ReleaseYunosTvpubadminContentChildLeafnodeGetAPIResponse 将 YunosTvpubadminContentChildLeafnodeGetAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentChildLeafnodeGetAPIResponse(v *YunosTvpubadminContentChildLeafnodeGetAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentChildLeafnodeGetAPIResponse.Put(v)
}
