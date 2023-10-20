package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugQuantityBatchUpdateAPIResponse 批量同步库存接口 API返回值
// taobao.drug.quantity.batch.update
//
// 商家通过top接口可以批量修改商品库存
type TaobaoDrugQuantityBatchUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoDrugQuantityBatchUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDrugQuantityBatchUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDrugQuantityBatchUpdateAPIResponseModel).Reset()
}

// TaobaoDrugQuantityBatchUpdateAPIResponseModel is 批量同步库存接口 成功返回结果
type TaobaoDrugQuantityBatchUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_quantity_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoDrugQuantityBatchUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDrugQuantityBatchUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDrugQuantityBatchUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDrugQuantityBatchUpdateAPIResponse)
	},
}

// GetTaobaoDrugQuantityBatchUpdateAPIResponse 从 sync.Pool 获取 TaobaoDrugQuantityBatchUpdateAPIResponse
func GetTaobaoDrugQuantityBatchUpdateAPIResponse() *TaobaoDrugQuantityBatchUpdateAPIResponse {
	return poolTaobaoDrugQuantityBatchUpdateAPIResponse.Get().(*TaobaoDrugQuantityBatchUpdateAPIResponse)
}

// ReleaseTaobaoDrugQuantityBatchUpdateAPIResponse 将 TaobaoDrugQuantityBatchUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDrugQuantityBatchUpdateAPIResponse(v *TaobaoDrugQuantityBatchUpdateAPIResponse) {
	v.Reset()
	poolTaobaoDrugQuantityBatchUpdateAPIResponse.Put(v)
}
