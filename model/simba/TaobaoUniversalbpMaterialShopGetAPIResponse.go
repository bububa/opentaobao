package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMaterialShopGetAPIResponse 获取店铺信息 API返回值
// taobao.universalbp.material.shop.get
//
// 获取店铺信息
type TaobaoUniversalbpMaterialShopGetAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpMaterialShopGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpMaterialShopGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpMaterialShopGetAPIResponseModel).Reset()
}

// TaobaoUniversalbpMaterialShopGetAPIResponseModel is 获取店铺信息 成功返回结果
type TaobaoUniversalbpMaterialShopGetAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_material_shop_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpMaterialShopGetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpMaterialShopGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpMaterialShopGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpMaterialShopGetAPIResponse)
	},
}

// GetTaobaoUniversalbpMaterialShopGetAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpMaterialShopGetAPIResponse
func GetTaobaoUniversalbpMaterialShopGetAPIResponse() *TaobaoUniversalbpMaterialShopGetAPIResponse {
	return poolTaobaoUniversalbpMaterialShopGetAPIResponse.Get().(*TaobaoUniversalbpMaterialShopGetAPIResponse)
}

// ReleaseTaobaoUniversalbpMaterialShopGetAPIResponse 将 TaobaoUniversalbpMaterialShopGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpMaterialShopGetAPIResponse(v *TaobaoUniversalbpMaterialShopGetAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpMaterialShopGetAPIResponse.Put(v)
}
