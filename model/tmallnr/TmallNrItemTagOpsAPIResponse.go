package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrItemTagOpsAPIResponse 区域零售商品打标去标 API返回值
// tmall.nr.item.tag.ops
//
// 参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
type TmallNrItemTagOpsAPIResponse struct {
	model.CommonResponse
	TmallNrItemTagOpsAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrItemTagOpsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrItemTagOpsAPIResponseModel).Reset()
}

// TmallNrItemTagOpsAPIResponseModel is 区域零售商品打标去标 成功返回结果
type TmallNrItemTagOpsAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_item_tag_ops_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrItemTagOpsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrItemTagOpsAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrItemTagOpsAPIResponse)
	},
}

// GetTmallNrItemTagOpsAPIResponse 从 sync.Pool 获取 TmallNrItemTagOpsAPIResponse
func GetTmallNrItemTagOpsAPIResponse() *TmallNrItemTagOpsAPIResponse {
	return poolTmallNrItemTagOpsAPIResponse.Get().(*TmallNrItemTagOpsAPIResponse)
}

// ReleaseTmallNrItemTagOpsAPIResponse 将 TmallNrItemTagOpsAPIResponse 保存到 sync.Pool
func ReleaseTmallNrItemTagOpsAPIResponse(v *TmallNrItemTagOpsAPIResponse) {
	v.Reset()
	poolTmallNrItemTagOpsAPIResponse.Put(v)
}
