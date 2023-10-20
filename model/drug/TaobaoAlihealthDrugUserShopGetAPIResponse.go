package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlihealthDrugUserShopGetAPIResponse 根据用户id获取店铺id API返回值
// taobao.alihealth.drug.user.shop.get
//
// 提供给千牛智能客服，获取用户当前咨询的店铺ID
type TaobaoAlihealthDrugUserShopGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlihealthDrugUserShopGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlihealthDrugUserShopGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlihealthDrugUserShopGetAPIResponseModel).Reset()
}

// TaobaoAlihealthDrugUserShopGetAPIResponseModel is 根据用户id获取店铺id 成功返回结果
type TaobaoAlihealthDrugUserShopGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alihealth_drug_user_shop_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// shopId
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlihealthDrugUserShopGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ShopId = 0
}

var poolTaobaoAlihealthDrugUserShopGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlihealthDrugUserShopGetAPIResponse)
	},
}

// GetTaobaoAlihealthDrugUserShopGetAPIResponse 从 sync.Pool 获取 TaobaoAlihealthDrugUserShopGetAPIResponse
func GetTaobaoAlihealthDrugUserShopGetAPIResponse() *TaobaoAlihealthDrugUserShopGetAPIResponse {
	return poolTaobaoAlihealthDrugUserShopGetAPIResponse.Get().(*TaobaoAlihealthDrugUserShopGetAPIResponse)
}

// ReleaseTaobaoAlihealthDrugUserShopGetAPIResponse 将 TaobaoAlihealthDrugUserShopGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlihealthDrugUserShopGetAPIResponse(v *TaobaoAlihealthDrugUserShopGetAPIResponse) {
	v.Reset()
	poolTaobaoAlihealthDrugUserShopGetAPIResponse.Put(v)
}
