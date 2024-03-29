package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemAddAPIResponse 新发商品 API返回值
// taobao.banamadpc.item.add
//
// 巴拿马供应商通过此接口新发商品
type TaobaoBanamadpcItemAddAPIResponse struct {
	model.CommonResponse
	TaobaoBanamadpcItemAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBanamadpcItemAddAPIResponseModel).Reset()
}

// TaobaoBanamadpcItemAddAPIResponseModel is 新发商品 成功返回结果
type TaobaoBanamadpcItemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TaobaoBanamadpcItemAddApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolTaobaoBanamadpcItemAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemAddAPIResponse)
	},
}

// GetTaobaoBanamadpcItemAddAPIResponse 从 sync.Pool 获取 TaobaoBanamadpcItemAddAPIResponse
func GetTaobaoBanamadpcItemAddAPIResponse() *TaobaoBanamadpcItemAddAPIResponse {
	return poolTaobaoBanamadpcItemAddAPIResponse.Get().(*TaobaoBanamadpcItemAddAPIResponse)
}

// ReleaseTaobaoBanamadpcItemAddAPIResponse 将 TaobaoBanamadpcItemAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBanamadpcItemAddAPIResponse(v *TaobaoBanamadpcItemAddAPIResponse) {
	v.Reset()
	poolTaobaoBanamadpcItemAddAPIResponse.Put(v)
}
