package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugPriceBatchUpdateAPIResponse 商家批量更新宝贝价格 API返回值
// taobao.drug.price.batch.update
//
// 商家批量更新宝贝价格
type TaobaoDrugPriceBatchUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoDrugPriceBatchUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDrugPriceBatchUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDrugPriceBatchUpdateAPIResponseModel).Reset()
}

// TaobaoDrugPriceBatchUpdateAPIResponseModel is 商家批量更新宝贝价格 成功返回结果
type TaobaoDrugPriceBatchUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_price_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoDrugPriceBatchUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDrugPriceBatchUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDrugPriceBatchUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDrugPriceBatchUpdateAPIResponse)
	},
}

// GetTaobaoDrugPriceBatchUpdateAPIResponse 从 sync.Pool 获取 TaobaoDrugPriceBatchUpdateAPIResponse
func GetTaobaoDrugPriceBatchUpdateAPIResponse() *TaobaoDrugPriceBatchUpdateAPIResponse {
	return poolTaobaoDrugPriceBatchUpdateAPIResponse.Get().(*TaobaoDrugPriceBatchUpdateAPIResponse)
}

// ReleaseTaobaoDrugPriceBatchUpdateAPIResponse 将 TaobaoDrugPriceBatchUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDrugPriceBatchUpdateAPIResponse(v *TaobaoDrugPriceBatchUpdateAPIResponse) {
	v.Reset()
	poolTaobaoDrugPriceBatchUpdateAPIResponse.Put(v)
}
