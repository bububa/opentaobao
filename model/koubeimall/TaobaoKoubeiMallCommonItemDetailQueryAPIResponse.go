package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonItemDetailQueryAPIResponse 查询商品详情信息 API返回值
// taobao.koubei.mall.common.item.detail.query
//
// 查询口碑综合体内商品详情信息
type TaobaoKoubeiMallCommonItemDetailQueryAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonItemDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel is 查询商品详情信息 成功返回结果
type TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_item_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiMallCommonItemDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonItemDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonItemDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonItemDetailQueryAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonItemDetailQueryAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonItemDetailQueryAPIResponse
func GetTaobaoKoubeiMallCommonItemDetailQueryAPIResponse() *TaobaoKoubeiMallCommonItemDetailQueryAPIResponse {
	return poolTaobaoKoubeiMallCommonItemDetailQueryAPIResponse.Get().(*TaobaoKoubeiMallCommonItemDetailQueryAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonItemDetailQueryAPIResponse 将 TaobaoKoubeiMallCommonItemDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonItemDetailQueryAPIResponse(v *TaobaoKoubeiMallCommonItemDetailQueryAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonItemDetailQueryAPIResponse.Put(v)
}
