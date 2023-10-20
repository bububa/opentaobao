package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreExtendUpdateAPIResponse 商户门店拓展信息更新接口 API返回值
// taobao.place.store.extend.update
//
// 更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
type TaobaoPlaceStoreExtendUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreExtendUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreExtendUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreExtendUpdateAPIResponseModel).Reset()
}

// TaobaoPlaceStoreExtendUpdateAPIResponseModel is 商户门店拓展信息更新接口 成功返回结果
type TaobaoPlaceStoreExtendUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_extend_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoPlaceStoreExtendUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreExtendUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStoreExtendUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreExtendUpdateAPIResponse)
	},
}

// GetTaobaoPlaceStoreExtendUpdateAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreExtendUpdateAPIResponse
func GetTaobaoPlaceStoreExtendUpdateAPIResponse() *TaobaoPlaceStoreExtendUpdateAPIResponse {
	return poolTaobaoPlaceStoreExtendUpdateAPIResponse.Get().(*TaobaoPlaceStoreExtendUpdateAPIResponse)
}

// ReleaseTaobaoPlaceStoreExtendUpdateAPIResponse 将 TaobaoPlaceStoreExtendUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreExtendUpdateAPIResponse(v *TaobaoPlaceStoreExtendUpdateAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreExtendUpdateAPIResponse.Put(v)
}
