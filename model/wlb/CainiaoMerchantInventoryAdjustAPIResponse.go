package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoMerchantInventoryAdjustAPIResponse 商家库存调整 API返回值
// cainiao.merchant.inventory.adjust
//
// 商家仓库存调整接口，目前仅支持全量更新
type CainiaoMerchantInventoryAdjustAPIResponse struct {
	model.CommonResponse
	CainiaoMerchantInventoryAdjustAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoMerchantInventoryAdjustAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoMerchantInventoryAdjustAPIResponseModel).Reset()
}

// CainiaoMerchantInventoryAdjustAPIResponseModel is 商家库存调整 成功返回结果
type CainiaoMerchantInventoryAdjustAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_merchant_inventory_adjust_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SingleResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoMerchantInventoryAdjustAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoMerchantInventoryAdjustAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoMerchantInventoryAdjustAPIResponse)
	},
}

// GetCainiaoMerchantInventoryAdjustAPIResponse 从 sync.Pool 获取 CainiaoMerchantInventoryAdjustAPIResponse
func GetCainiaoMerchantInventoryAdjustAPIResponse() *CainiaoMerchantInventoryAdjustAPIResponse {
	return poolCainiaoMerchantInventoryAdjustAPIResponse.Get().(*CainiaoMerchantInventoryAdjustAPIResponse)
}

// ReleaseCainiaoMerchantInventoryAdjustAPIResponse 将 CainiaoMerchantInventoryAdjustAPIResponse 保存到 sync.Pool
func ReleaseCainiaoMerchantInventoryAdjustAPIResponse(v *CainiaoMerchantInventoryAdjustAPIResponse) {
	v.Reset()
	poolCainiaoMerchantInventoryAdjustAPIResponse.Put(v)
}
