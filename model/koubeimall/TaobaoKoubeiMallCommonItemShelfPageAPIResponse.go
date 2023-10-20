package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonItemShelfPageAPIResponse 门店货架商品列表信息查询 API返回值
// taobao.koubei.mall.common.item.shelf.page
//
// 查询口碑综合体内门店货架商品列表信息接口
type TaobaoKoubeiMallCommonItemShelfPageAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonItemShelfPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonItemShelfPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonItemShelfPageAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonItemShelfPageAPIResponseModel is 门店货架商品列表信息查询 成功返回结果
type TaobaoKoubeiMallCommonItemShelfPageAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_item_shelf_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// API接口返回的result模型
	Result *TaobaoKoubeiMallCommonItemShelfPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonItemShelfPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonItemShelfPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonItemShelfPageAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonItemShelfPageAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonItemShelfPageAPIResponse
func GetTaobaoKoubeiMallCommonItemShelfPageAPIResponse() *TaobaoKoubeiMallCommonItemShelfPageAPIResponse {
	return poolTaobaoKoubeiMallCommonItemShelfPageAPIResponse.Get().(*TaobaoKoubeiMallCommonItemShelfPageAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonItemShelfPageAPIResponse 将 TaobaoKoubeiMallCommonItemShelfPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonItemShelfPageAPIResponse(v *TaobaoKoubeiMallCommonItemShelfPageAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonItemShelfPageAPIResponse.Put(v)
}
