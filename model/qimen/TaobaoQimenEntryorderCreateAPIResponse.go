package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderCreateAPIResponse 入库单创建接口 API返回值
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
type TaobaoQimenEntryorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenEntryorderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenEntryorderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenEntryorderCreateAPIResponseModel).Reset()
}

// TaobaoQimenEntryorderCreateAPIResponseModel is 入库单创建接口 成功返回结果
type TaobaoQimenEntryorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_entryorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenEntryorderCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenEntryorderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenEntryorderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderCreateAPIResponse)
	},
}

// GetTaobaoQimenEntryorderCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenEntryorderCreateAPIResponse
func GetTaobaoQimenEntryorderCreateAPIResponse() *TaobaoQimenEntryorderCreateAPIResponse {
	return poolTaobaoQimenEntryorderCreateAPIResponse.Get().(*TaobaoQimenEntryorderCreateAPIResponse)
}

// ReleaseTaobaoQimenEntryorderCreateAPIResponse 将 TaobaoQimenEntryorderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenEntryorderCreateAPIResponse(v *TaobaoQimenEntryorderCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenEntryorderCreateAPIResponse.Put(v)
}
