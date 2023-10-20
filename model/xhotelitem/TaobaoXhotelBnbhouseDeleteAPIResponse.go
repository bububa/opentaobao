package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbhouseDeleteAPIResponse 民宿门店删除接口 API返回值
// taobao.xhotel.bnbhouse.delete
//
// 支持门店相关的门店删除，删除门店会级联删除门店下面的房源
type TaobaoXhotelBnbhouseDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbhouseDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbhouseDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbhouseDeleteAPIResponseModel).Reset()
}

// TaobaoXhotelBnbhouseDeleteAPIResponseModel is 民宿门店删除接口 成功返回结果
type TaobaoXhotelBnbhouseDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbhouse_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否出错
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否出错
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbhouseDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Error = false
}

var poolTaobaoXhotelBnbhouseDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbhouseDeleteAPIResponse)
	},
}

// GetTaobaoXhotelBnbhouseDeleteAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbhouseDeleteAPIResponse
func GetTaobaoXhotelBnbhouseDeleteAPIResponse() *TaobaoXhotelBnbhouseDeleteAPIResponse {
	return poolTaobaoXhotelBnbhouseDeleteAPIResponse.Get().(*TaobaoXhotelBnbhouseDeleteAPIResponse)
}

// ReleaseTaobaoXhotelBnbhouseDeleteAPIResponse 将 TaobaoXhotelBnbhouseDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbhouseDeleteAPIResponse(v *TaobaoXhotelBnbhouseDeleteAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbhouseDeleteAPIResponse.Put(v)
}
