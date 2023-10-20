package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse 查询门店推荐菜信息 API返回值
// taobao.koubei.mall.common.store.display.goods.list
//
// 提供查询口碑商圈内的门店推荐菜信息
type TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponseModel is 查询门店推荐菜信息 成功返回结果
type TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_display_goods_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// API接口返回的result模型
	Result *TaobaoKoubeiMallCommonStoreDisplayGoodsListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse
func GetTaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse() *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse {
	return poolTaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse.Get().(*TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse 将 TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse(v *TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse.Put(v)
}
