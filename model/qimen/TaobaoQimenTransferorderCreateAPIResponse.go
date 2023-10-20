package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTransferorderCreateAPIResponse 调拨单创建 API返回值
// taobao.qimen.transferorder.create
//
// 调拨单创建
type TaobaoQimenTransferorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTransferorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenTransferorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenTransferorderCreateAPIResponseModel).Reset()
}

// TaobaoQimenTransferorderCreateAPIResponseModel is 调拨单创建 成功返回结果
type TaobaoQimenTransferorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_transferorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenTransferorderCreateStruct `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenTransferorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenTransferorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenTransferorderCreateAPIResponse)
	},
}

// GetTaobaoQimenTransferorderCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenTransferorderCreateAPIResponse
func GetTaobaoQimenTransferorderCreateAPIResponse() *TaobaoQimenTransferorderCreateAPIResponse {
	return poolTaobaoQimenTransferorderCreateAPIResponse.Get().(*TaobaoQimenTransferorderCreateAPIResponse)
}

// ReleaseTaobaoQimenTransferorderCreateAPIResponse 将 TaobaoQimenTransferorderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenTransferorderCreateAPIResponse(v *TaobaoQimenTransferorderCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenTransferorderCreateAPIResponse.Put(v)
}
