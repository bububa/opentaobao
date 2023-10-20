package tuanhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelComboOffshelfAPIResponse 酒店套餐下架 API返回值
// taobao.xhotel.combo.offshelf
//
// 酒店套餐下架
type TaobaoXhotelComboOffshelfAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelComboOffshelfAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelComboOffshelfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelComboOffshelfAPIResponseModel).Reset()
}

// TaobaoXhotelComboOffshelfAPIResponseModel is 酒店套餐下架 成功返回结果
type TaobaoXhotelComboOffshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_combo_offshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下架状态
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelComboOffshelfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolTaobaoXhotelComboOffshelfAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelComboOffshelfAPIResponse)
	},
}

// GetTaobaoXhotelComboOffshelfAPIResponse 从 sync.Pool 获取 TaobaoXhotelComboOffshelfAPIResponse
func GetTaobaoXhotelComboOffshelfAPIResponse() *TaobaoXhotelComboOffshelfAPIResponse {
	return poolTaobaoXhotelComboOffshelfAPIResponse.Get().(*TaobaoXhotelComboOffshelfAPIResponse)
}

// ReleaseTaobaoXhotelComboOffshelfAPIResponse 将 TaobaoXhotelComboOffshelfAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelComboOffshelfAPIResponse(v *TaobaoXhotelComboOffshelfAPIResponse) {
	v.Reset()
	poolTaobaoXhotelComboOffshelfAPIResponse.Put(v)
}
