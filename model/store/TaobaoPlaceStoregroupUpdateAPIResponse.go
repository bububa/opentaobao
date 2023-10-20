package store

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoregroupUpdateAPIResponse 门店库修改基本信息 API返回值
// taobao.place.storegroup.update
//
// 门店库修改基本信息
type TaobaoPlaceStoregroupUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoregroupUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoregroupUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoregroupUpdateAPIResponseModel).Reset()
}

// TaobaoPlaceStoregroupUpdateAPIResponseModel is 门店库修改基本信息 成功返回结果
type TaobaoPlaceStoregroupUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"place_storegroup_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoregroupUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoPlaceStoregroupUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoregroupUpdateAPIResponse)
	},
}

// GetTaobaoPlaceStoregroupUpdateAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoregroupUpdateAPIResponse
func GetTaobaoPlaceStoregroupUpdateAPIResponse() *TaobaoPlaceStoregroupUpdateAPIResponse {
	return poolTaobaoPlaceStoregroupUpdateAPIResponse.Get().(*TaobaoPlaceStoregroupUpdateAPIResponse)
}

// ReleaseTaobaoPlaceStoregroupUpdateAPIResponse 将 TaobaoPlaceStoregroupUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoregroupUpdateAPIResponse(v *TaobaoPlaceStoregroupUpdateAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoregroupUpdateAPIResponse.Put(v)
}
