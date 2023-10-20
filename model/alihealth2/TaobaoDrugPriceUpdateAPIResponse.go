package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugPriceUpdateAPIResponse 商家更新宝贝价格 API返回值
// taobao.drug.price.update
//
// 商家更新价格
type TaobaoDrugPriceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoDrugPriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDrugPriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDrugPriceUpdateAPIResponseModel).Reset()
}

// TaobaoDrugPriceUpdateAPIResponseModel is 商家更新宝贝价格 成功返回结果
type TaobaoDrugPriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_price_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoDrugPriceUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDrugPriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDrugPriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDrugPriceUpdateAPIResponse)
	},
}

// GetTaobaoDrugPriceUpdateAPIResponse 从 sync.Pool 获取 TaobaoDrugPriceUpdateAPIResponse
func GetTaobaoDrugPriceUpdateAPIResponse() *TaobaoDrugPriceUpdateAPIResponse {
	return poolTaobaoDrugPriceUpdateAPIResponse.Get().(*TaobaoDrugPriceUpdateAPIResponse)
}

// ReleaseTaobaoDrugPriceUpdateAPIResponse 将 TaobaoDrugPriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDrugPriceUpdateAPIResponse(v *TaobaoDrugPriceUpdateAPIResponse) {
	v.Reset()
	poolTaobaoDrugPriceUpdateAPIResponse.Put(v)
}
