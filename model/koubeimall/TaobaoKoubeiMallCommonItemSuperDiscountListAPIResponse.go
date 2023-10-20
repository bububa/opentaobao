package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse 查询商圈内的超值特惠商品信息 API返回值
// taobao.koubei.mall.common.item.super.discount.list
//
// 查询商圈超值特惠商品信息列表
type TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponseModel is 查询商圈内的超值特惠商品信息 成功返回结果
type TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_item_super_discount_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiMallCommonItemSuperDiscountListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse
func GetTaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse() *TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse {
	return poolTaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse.Get().(*TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse 将 TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse(v *TaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonItemSuperDiscountListAPIResponse.Put(v)
}
