package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderQueryAPIResponse 入库单查询接口 API返回值
// taobao.qimen.entryorder.query
//
// ERP调用接口，查询入库单信息;
type TaobaoQimenEntryorderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenEntryorderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenEntryorderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenEntryorderQueryAPIResponseModel).Reset()
}

// TaobaoQimenEntryorderQueryAPIResponseModel is 入库单查询接口 成功返回结果
type TaobaoQimenEntryorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_entryorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *EntryOrderQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenEntryorderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenEntryorderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderQueryAPIResponse)
	},
}

// GetTaobaoQimenEntryorderQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenEntryorderQueryAPIResponse
func GetTaobaoQimenEntryorderQueryAPIResponse() *TaobaoQimenEntryorderQueryAPIResponse {
	return poolTaobaoQimenEntryorderQueryAPIResponse.Get().(*TaobaoQimenEntryorderQueryAPIResponse)
}

// ReleaseTaobaoQimenEntryorderQueryAPIResponse 将 TaobaoQimenEntryorderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenEntryorderQueryAPIResponse(v *TaobaoQimenEntryorderQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenEntryorderQueryAPIResponse.Put(v)
}
