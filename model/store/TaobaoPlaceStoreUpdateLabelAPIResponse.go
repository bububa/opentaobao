package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreUpdateLabelAPIResponse 商户门店标签更新接口 API返回值
// taobao.place.store.update.label
//
// 更新商户门店标签（服务、权益、标签）接口
type TaobaoPlaceStoreUpdateLabelAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreUpdateLabelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreUpdateLabelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreUpdateLabelAPIResponseModel).Reset()
}

// TaobaoPlaceStoreUpdateLabelAPIResponseModel is 商户门店标签更新接口 成功返回结果
type TaobaoPlaceStoreUpdateLabelAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_update_label_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *UpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreUpdateLabelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStoreUpdateLabelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreUpdateLabelAPIResponse)
	},
}

// GetTaobaoPlaceStoreUpdateLabelAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreUpdateLabelAPIResponse
func GetTaobaoPlaceStoreUpdateLabelAPIResponse() *TaobaoPlaceStoreUpdateLabelAPIResponse {
	return poolTaobaoPlaceStoreUpdateLabelAPIResponse.Get().(*TaobaoPlaceStoreUpdateLabelAPIResponse)
}

// ReleaseTaobaoPlaceStoreUpdateLabelAPIResponse 将 TaobaoPlaceStoreUpdateLabelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreUpdateLabelAPIResponse(v *TaobaoPlaceStoreUpdateLabelAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreUpdateLabelAPIResponse.Put(v)
}
