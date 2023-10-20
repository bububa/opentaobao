package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreprocessConfirmAPIResponse 仓内加工单确认接口 API返回值
// taobao.qimen.storeprocess.confirm
//
// WMS调用奇门的接口,回传仓内加工单创建情况
type TaobaoQimenStoreprocessConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStoreprocessConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStoreprocessConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStoreprocessConfirmAPIResponseModel).Reset()
}

// TaobaoQimenStoreprocessConfirmAPIResponseModel is 仓内加工单确认接口 成功返回结果
type TaobaoQimenStoreprocessConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_storeprocess_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenStoreprocessConfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStoreprocessConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenStoreprocessConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStoreprocessConfirmAPIResponse)
	},
}

// GetTaobaoQimenStoreprocessConfirmAPIResponse 从 sync.Pool 获取 TaobaoQimenStoreprocessConfirmAPIResponse
func GetTaobaoQimenStoreprocessConfirmAPIResponse() *TaobaoQimenStoreprocessConfirmAPIResponse {
	return poolTaobaoQimenStoreprocessConfirmAPIResponse.Get().(*TaobaoQimenStoreprocessConfirmAPIResponse)
}

// ReleaseTaobaoQimenStoreprocessConfirmAPIResponse 将 TaobaoQimenStoreprocessConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStoreprocessConfirmAPIResponse(v *TaobaoQimenStoreprocessConfirmAPIResponse) {
	v.Reset()
	poolTaobaoQimenStoreprocessConfirmAPIResponse.Put(v)
}
