package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse 查询综合体内的门店详细信息 API返回值
// taobao.koubei.mall.common.store.detail.query
//
// 查询口碑综合体内的门店详情信息
type TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonStoreDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonStoreDetailQueryAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonStoreDetailQueryAPIResponseModel is 查询综合体内的门店详细信息 成功返回结果
type TaobaoKoubeiMallCommonStoreDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoKoubeiMallCommonStoreDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStoreDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonStoreDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonStoreDetailQueryAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse
func GetTaobaoKoubeiMallCommonStoreDetailQueryAPIResponse() *TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse {
	return poolTaobaoKoubeiMallCommonStoreDetailQueryAPIResponse.Get().(*TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonStoreDetailQueryAPIResponse 将 TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonStoreDetailQueryAPIResponse(v *TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonStoreDetailQueryAPIResponse.Put(v)
}
