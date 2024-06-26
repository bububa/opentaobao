package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSimpleschemaAddAPIResponse 天猫简化发布商品 API返回值
// tmall.item.simpleschema.add
//
// 天猫简化版schema发布商品。
type TmallItemSimpleschemaAddAPIResponse struct {
	model.CommonResponse
	TmallItemSimpleschemaAddAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSimpleschemaAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSimpleschemaAddAPIResponseModel).Reset()
}

// TmallItemSimpleschemaAddAPIResponseModel is 天猫简化发布商品 成功返回结果
type TmallItemSimpleschemaAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_simpleschema_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发布成功后返回商品ID
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 商品最后发布时间。
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSimpleschemaAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.GmtModified = ""
}

var poolTmallItemSimpleschemaAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSimpleschemaAddAPIResponse)
	},
}

// GetTmallItemSimpleschemaAddAPIResponse 从 sync.Pool 获取 TmallItemSimpleschemaAddAPIResponse
func GetTmallItemSimpleschemaAddAPIResponse() *TmallItemSimpleschemaAddAPIResponse {
	return poolTmallItemSimpleschemaAddAPIResponse.Get().(*TmallItemSimpleschemaAddAPIResponse)
}

// ReleaseTmallItemSimpleschemaAddAPIResponse 将 TmallItemSimpleschemaAddAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSimpleschemaAddAPIResponse(v *TmallItemSimpleschemaAddAPIResponse) {
	v.Reset()
	poolTmallItemSimpleschemaAddAPIResponse.Put(v)
}
