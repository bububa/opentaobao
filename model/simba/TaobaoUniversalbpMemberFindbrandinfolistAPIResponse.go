package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMemberFindbrandinfolistAPIResponse 查询可用品牌列表 API返回值
// taobao.universalbp.member.findbrandinfolist
//
// 查询账号对应的品牌，用于品牌人群屏蔽等
type TaobaoUniversalbpMemberFindbrandinfolistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpMemberFindbrandinfolistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpMemberFindbrandinfolistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpMemberFindbrandinfolistAPIResponseModel).Reset()
}

// TaobaoUniversalbpMemberFindbrandinfolistAPIResponseModel is 查询可用品牌列表 成功返回结果
type TaobaoUniversalbpMemberFindbrandinfolistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_member_findbrandinfolist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpMemberFindbrandinfolistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpMemberFindbrandinfolistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpMemberFindbrandinfolistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpMemberFindbrandinfolistAPIResponse)
	},
}

// GetTaobaoUniversalbpMemberFindbrandinfolistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpMemberFindbrandinfolistAPIResponse
func GetTaobaoUniversalbpMemberFindbrandinfolistAPIResponse() *TaobaoUniversalbpMemberFindbrandinfolistAPIResponse {
	return poolTaobaoUniversalbpMemberFindbrandinfolistAPIResponse.Get().(*TaobaoUniversalbpMemberFindbrandinfolistAPIResponse)
}

// ReleaseTaobaoUniversalbpMemberFindbrandinfolistAPIResponse 将 TaobaoUniversalbpMemberFindbrandinfolistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpMemberFindbrandinfolistAPIResponse(v *TaobaoUniversalbpMemberFindbrandinfolistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpMemberFindbrandinfolistAPIResponse.Put(v)
}
