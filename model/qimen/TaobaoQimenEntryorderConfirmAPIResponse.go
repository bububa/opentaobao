package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderConfirmAPIResponse 入库单确认接口 API返回值
// taobao.qimen.entryorder.confirm
//
// WMS调用接口，回传入库单信息;
type TaobaoQimenEntryorderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoQimenEntryorderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenEntryorderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenEntryorderConfirmAPIResponseModel).Reset()
}

// TaobaoQimenEntryorderConfirmAPIResponseModel is 入库单确认接口 成功返回结果
type TaobaoQimenEntryorderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_entryorder_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenEntryorderConfirmResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenEntryorderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenEntryorderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderConfirmAPIResponse)
	},
}

// GetTaobaoQimenEntryorderConfirmAPIResponse 从 sync.Pool 获取 TaobaoQimenEntryorderConfirmAPIResponse
func GetTaobaoQimenEntryorderConfirmAPIResponse() *TaobaoQimenEntryorderConfirmAPIResponse {
	return poolTaobaoQimenEntryorderConfirmAPIResponse.Get().(*TaobaoQimenEntryorderConfirmAPIResponse)
}

// ReleaseTaobaoQimenEntryorderConfirmAPIResponse 将 TaobaoQimenEntryorderConfirmAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenEntryorderConfirmAPIResponse(v *TaobaoQimenEntryorderConfirmAPIResponse) {
	v.Reset()
	poolTaobaoQimenEntryorderConfirmAPIResponse.Put(v)
}
