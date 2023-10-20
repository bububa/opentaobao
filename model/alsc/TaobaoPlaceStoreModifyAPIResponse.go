package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreModifyAPIResponse 商家修改线下门店 API返回值
// taobao.place.store.modify
//
// 用于商家修改线下门店信息
type TaobaoPlaceStoreModifyAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreModifyAPIResponseModel).Reset()
}

// TaobaoPlaceStoreModifyAPIResponseModel is 商家修改线下门店 成功返回结果
type TaobaoPlaceStoreModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否修改成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoPlaceStoreModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreModifyAPIResponse)
	},
}

// GetTaobaoPlaceStoreModifyAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreModifyAPIResponse
func GetTaobaoPlaceStoreModifyAPIResponse() *TaobaoPlaceStoreModifyAPIResponse {
	return poolTaobaoPlaceStoreModifyAPIResponse.Get().(*TaobaoPlaceStoreModifyAPIResponse)
}

// ReleaseTaobaoPlaceStoreModifyAPIResponse 将 TaobaoPlaceStoreModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreModifyAPIResponse(v *TaobaoPlaceStoreModifyAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreModifyAPIResponse.Put(v)
}
