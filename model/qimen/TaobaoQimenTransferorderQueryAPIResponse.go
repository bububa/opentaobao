package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTransferorderQueryAPIResponse 调拨单查询 API返回值
// taobao.qimen.transferorder.query
//
// 调拨单查询
type TaobaoQimenTransferorderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTransferorderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenTransferorderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenTransferorderQueryAPIResponseModel).Reset()
}

// TaobaoQimenTransferorderQueryAPIResponseModel is 调拨单查询 成功返回结果
type TaobaoQimenTransferorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_transferorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenTransferorderQueryStruct `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenTransferorderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenTransferorderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenTransferorderQueryAPIResponse)
	},
}

// GetTaobaoQimenTransferorderQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenTransferorderQueryAPIResponse
func GetTaobaoQimenTransferorderQueryAPIResponse() *TaobaoQimenTransferorderQueryAPIResponse {
	return poolTaobaoQimenTransferorderQueryAPIResponse.Get().(*TaobaoQimenTransferorderQueryAPIResponse)
}

// ReleaseTaobaoQimenTransferorderQueryAPIResponse 将 TaobaoQimenTransferorderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenTransferorderQueryAPIResponse(v *TaobaoQimenTransferorderQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenTransferorderQueryAPIResponse.Put(v)
}
