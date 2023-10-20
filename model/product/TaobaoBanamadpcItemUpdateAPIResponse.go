package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemUpdateAPIResponse 编辑商品 API返回值
// taobao.banamadpc.item.update
//
// 巴拿马供应商通过此接口编辑商品
type TaobaoBanamadpcItemUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoBanamadpcItemUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBanamadpcItemUpdateAPIResponseModel).Reset()
}

// TaobaoBanamadpcItemUpdateAPIResponseModel is 编辑商品 成功返回结果
type TaobaoBanamadpcItemUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TaobaoBanamadpcItemUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolTaobaoBanamadpcItemUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemUpdateAPIResponse)
	},
}

// GetTaobaoBanamadpcItemUpdateAPIResponse 从 sync.Pool 获取 TaobaoBanamadpcItemUpdateAPIResponse
func GetTaobaoBanamadpcItemUpdateAPIResponse() *TaobaoBanamadpcItemUpdateAPIResponse {
	return poolTaobaoBanamadpcItemUpdateAPIResponse.Get().(*TaobaoBanamadpcItemUpdateAPIResponse)
}

// ReleaseTaobaoBanamadpcItemUpdateAPIResponse 将 TaobaoBanamadpcItemUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBanamadpcItemUpdateAPIResponse(v *TaobaoBanamadpcItemUpdateAPIResponse) {
	v.Reset()
	poolTaobaoBanamadpcItemUpdateAPIResponse.Put(v)
}
