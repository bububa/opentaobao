package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrInventoryUpdateAPIResponse 门店业务同步库存 API返回值
// tmall.nr.inventory.update
//
// 用于商家每日同步更新门店库存
type TmallNrInventoryUpdateAPIResponse struct {
	model.CommonResponse
	TmallNrInventoryUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrInventoryUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrInventoryUpdateAPIResponseModel).Reset()
}

// TmallNrInventoryUpdateAPIResponseModel is 门店业务同步库存 成功返回结果
type TmallNrInventoryUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_inventory_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrInventoryUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrInventoryUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrInventoryUpdateAPIResponse)
	},
}

// GetTmallNrInventoryUpdateAPIResponse 从 sync.Pool 获取 TmallNrInventoryUpdateAPIResponse
func GetTmallNrInventoryUpdateAPIResponse() *TmallNrInventoryUpdateAPIResponse {
	return poolTmallNrInventoryUpdateAPIResponse.Get().(*TmallNrInventoryUpdateAPIResponse)
}

// ReleaseTmallNrInventoryUpdateAPIResponse 将 TmallNrInventoryUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallNrInventoryUpdateAPIResponse(v *TmallNrInventoryUpdateAPIResponse) {
	v.Reset()
	poolTmallNrInventoryUpdateAPIResponse.Put(v)
}
