package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnorderConfirmAPIResponse 退货入库单确认接口 API返回值
// taobao.qimen.returnorder.confirm
//
// taobao.qimen.returnorder.confirm
type TaobaoQimenReturnorderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenReturnorderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenReturnorderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenReturnorderConfirmAPIResponseModel).Reset()
}

// TaobaoQimenReturnorderConfirmAPIResponseModel is 退货入库单确认接口 成功返回结果
type TaobaoQimenReturnorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_returnorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenReturnorderConfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenReturnorderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenReturnorderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnorderConfirmAPIResponse)
	},
}

// GetTaobaoQimenReturnorderConfirmAPIResponse 从 sync.Pool 获取 TaobaoQimenReturnorderConfirmAPIResponse
func GetTaobaoQimenReturnorderConfirmAPIResponse() *TaobaoQimenReturnorderConfirmAPIResponse {
	return poolTaobaoQimenReturnorderConfirmAPIResponse.Get().(*TaobaoQimenReturnorderConfirmAPIResponse)
}

// ReleaseTaobaoQimenReturnorderConfirmAPIResponse 将 TaobaoQimenReturnorderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenReturnorderConfirmAPIResponse(v *TaobaoQimenReturnorderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoQimenReturnorderConfirmAPIResponse.Put(v)
}
