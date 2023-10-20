package koubeimall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiMallCommonStorePageAPIResponse 分页查询综合体内的门店列表信息 API返回值
// taobao.koubei.mall.common.store.page
//
// 分页查询综合体内的门店列表信息
type TaobaoKoubeiMallCommonStorePageAPIResponse struct {
	model.CommonResponse
	TaobaoKoubeiMallCommonStorePageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStorePageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKoubeiMallCommonStorePageAPIResponseModel).Reset()
}

// TaobaoKoubeiMallCommonStorePageAPIResponseModel is 分页查询综合体内的门店列表信息 成功返回结果
type TaobaoKoubeiMallCommonStorePageAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// API接口返回的result模型
	Result *TaobaoKoubeiMallCommonStorePageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKoubeiMallCommonStorePageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoKoubeiMallCommonStorePageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonStorePageAPIResponse)
	},
}

// GetTaobaoKoubeiMallCommonStorePageAPIResponse 从 sync.Pool 获取 TaobaoKoubeiMallCommonStorePageAPIResponse
func GetTaobaoKoubeiMallCommonStorePageAPIResponse() *TaobaoKoubeiMallCommonStorePageAPIResponse {
	return poolTaobaoKoubeiMallCommonStorePageAPIResponse.Get().(*TaobaoKoubeiMallCommonStorePageAPIResponse)
}

// ReleaseTaobaoKoubeiMallCommonStorePageAPIResponse 将 TaobaoKoubeiMallCommonStorePageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKoubeiMallCommonStorePageAPIResponse(v *TaobaoKoubeiMallCommonStorePageAPIResponse) {
	v.Reset()
	poolTaobaoKoubeiMallCommonStorePageAPIResponse.Put(v)
}
