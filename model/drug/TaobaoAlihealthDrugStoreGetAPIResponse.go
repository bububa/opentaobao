package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlihealthDrugStoreGetAPIResponse 根据店铺id获取店铺详情 API返回值
// taobao.alihealth.drug.store.get
//
// 根据店铺id获取店铺详情
type TaobaoAlihealthDrugStoreGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlihealthDrugStoreGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlihealthDrugStoreGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlihealthDrugStoreGetAPIResponseModel).Reset()
}

// TaobaoAlihealthDrugStoreGetAPIResponseModel is 根据店铺id获取店铺详情 成功返回结果
type TaobaoAlihealthDrugStoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alihealth_drug_store_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model *StoreDetailDto `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlihealthDrugStoreGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = nil
}

var poolTaobaoAlihealthDrugStoreGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlihealthDrugStoreGetAPIResponse)
	},
}

// GetTaobaoAlihealthDrugStoreGetAPIResponse 从 sync.Pool 获取 TaobaoAlihealthDrugStoreGetAPIResponse
func GetTaobaoAlihealthDrugStoreGetAPIResponse() *TaobaoAlihealthDrugStoreGetAPIResponse {
	return poolTaobaoAlihealthDrugStoreGetAPIResponse.Get().(*TaobaoAlihealthDrugStoreGetAPIResponse)
}

// ReleaseTaobaoAlihealthDrugStoreGetAPIResponse 将 TaobaoAlihealthDrugStoreGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlihealthDrugStoreGetAPIResponse(v *TaobaoAlihealthDrugStoreGetAPIResponse) {
	v.Reset()
	poolTaobaoAlihealthDrugStoreGetAPIResponse.Put(v)
}
