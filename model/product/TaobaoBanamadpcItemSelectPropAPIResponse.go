package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemSelectPropAPIResponse 获取子属性 API返回值
// taobao.banamadpc.item.select.prop
//
// 巴拿马供应商通过此接口获取子属性
type TaobaoBanamadpcItemSelectPropAPIResponse struct {
	model.CommonResponse
	TaobaoBanamadpcItemSelectPropAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemSelectPropAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBanamadpcItemSelectPropAPIResponseModel).Reset()
}

// TaobaoBanamadpcItemSelectPropAPIResponseModel is 获取子属性 成功返回结果
type TaobaoBanamadpcItemSelectPropAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_select_prop_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TaobaoBanamadpcItemSelectPropApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBanamadpcItemSelectPropAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolTaobaoBanamadpcItemSelectPropAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemSelectPropAPIResponse)
	},
}

// GetTaobaoBanamadpcItemSelectPropAPIResponse 从 sync.Pool 获取 TaobaoBanamadpcItemSelectPropAPIResponse
func GetTaobaoBanamadpcItemSelectPropAPIResponse() *TaobaoBanamadpcItemSelectPropAPIResponse {
	return poolTaobaoBanamadpcItemSelectPropAPIResponse.Get().(*TaobaoBanamadpcItemSelectPropAPIResponse)
}

// ReleaseTaobaoBanamadpcItemSelectPropAPIResponse 将 TaobaoBanamadpcItemSelectPropAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBanamadpcItemSelectPropAPIResponse(v *TaobaoBanamadpcItemSelectPropAPIResponse) {
	v.Reset()
	poolTaobaoBanamadpcItemSelectPropAPIResponse.Put(v)
}
