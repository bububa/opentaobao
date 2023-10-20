package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnorderCreateAPIResponse 退货入库单创建接口 API返回值
// taobao.qimen.returnorder.create
//
// taobao.qimen.returnorder.create
type TaobaoQimenReturnorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenReturnorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenReturnorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenReturnorderCreateAPIResponseModel).Reset()
}

// TaobaoQimenReturnorderCreateAPIResponseModel is 退货入库单创建接口 成功返回结果
type TaobaoQimenReturnorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_returnorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenReturnorderCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenReturnorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenReturnorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnorderCreateAPIResponse)
	},
}

// GetTaobaoQimenReturnorderCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenReturnorderCreateAPIResponse
func GetTaobaoQimenReturnorderCreateAPIResponse() *TaobaoQimenReturnorderCreateAPIResponse {
	return poolTaobaoQimenReturnorderCreateAPIResponse.Get().(*TaobaoQimenReturnorderCreateAPIResponse)
}

// ReleaseTaobaoQimenReturnorderCreateAPIResponse 将 TaobaoQimenReturnorderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenReturnorderCreateAPIResponse(v *TaobaoQimenReturnorderCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenReturnorderCreateAPIResponse.Put(v)
}
