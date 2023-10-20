package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDrugShopListAPIResponse 查询卖家外卖店列表 API返回值
// taobao.drug.shop.list
//
// 查询卖家外卖店列表
type TaobaoDrugShopListAPIResponse struct {
	model.CommonResponse
	TaobaoDrugShopListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDrugShopListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDrugShopListAPIResponseModel).Reset()
}

// TaobaoDrugShopListAPIResponseModel is 查询卖家外卖店列表 成功返回结果
type TaobaoDrugShopListAPIResponseModel struct {
	XMLName xml.Name `xml:"drug_shop_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据结果集
	Result *TakeoutShopPage `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDrugShopListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDrugShopListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDrugShopListAPIResponse)
	},
}

// GetTaobaoDrugShopListAPIResponse 从 sync.Pool 获取 TaobaoDrugShopListAPIResponse
func GetTaobaoDrugShopListAPIResponse() *TaobaoDrugShopListAPIResponse {
	return poolTaobaoDrugShopListAPIResponse.Get().(*TaobaoDrugShopListAPIResponse)
}

// ReleaseTaobaoDrugShopListAPIResponse 将 TaobaoDrugShopListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDrugShopListAPIResponse(v *TaobaoDrugShopListAPIResponse) {
	v.Reset()
	poolTaobaoDrugShopListAPIResponse.Put(v)
}
