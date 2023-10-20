package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryAdjustExternalAPIResponse 非交易库存调整单 API返回值
// taobao.inventory.adjust.external
//
// 商家非交易调整库存，调拨出库、盘点等时调用
type TaobaoInventoryAdjustExternalAPIResponse struct {
	model.CommonResponse
	TaobaoInventoryAdjustExternalAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoInventoryAdjustExternalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoInventoryAdjustExternalAPIResponseModel).Reset()
}

// TaobaoInventoryAdjustExternalAPIResponseModel is 非交易库存调整单 成功返回结果
type TaobaoInventoryAdjustExternalAPIResponseModel struct {
	XMLName xml.Name `xml:"inventory_adjust_external_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提示信息
	TipInfos []TipInfo `json:"tip_infos,omitempty" xml:"tip_infos>tip_info,omitempty"`
	// 操作返回码
	OperateCode string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoInventoryAdjustExternalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TipInfos = m.TipInfos[:0]
	m.OperateCode = ""
}

var poolTaobaoInventoryAdjustExternalAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryAdjustExternalAPIResponse)
	},
}

// GetTaobaoInventoryAdjustExternalAPIResponse 从 sync.Pool 获取 TaobaoInventoryAdjustExternalAPIResponse
func GetTaobaoInventoryAdjustExternalAPIResponse() *TaobaoInventoryAdjustExternalAPIResponse {
	return poolTaobaoInventoryAdjustExternalAPIResponse.Get().(*TaobaoInventoryAdjustExternalAPIResponse)
}

// ReleaseTaobaoInventoryAdjustExternalAPIResponse 将 TaobaoInventoryAdjustExternalAPIResponse 保存到 sync.Pool
func ReleaseTaobaoInventoryAdjustExternalAPIResponse(v *TaobaoInventoryAdjustExternalAPIResponse) {
	v.Reset()
	poolTaobaoInventoryAdjustExternalAPIResponse.Put(v)
}
