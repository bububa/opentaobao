package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbhouseAddAPIResponse 民宿门店信息添加 API返回值
// taobao.xhotel.bnbhouse.add
//
// 添加和更新民宿门店的信息
type TaobaoXhotelBnbhouseAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbhouseAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbhouseAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbhouseAddAPIResponseModel).Reset()
}

// TaobaoXhotelBnbhouseAddAPIResponseModel is 民宿门店信息添加 成功返回结果
type TaobaoXhotelBnbhouseAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbhouse_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Results []XHotel `json:"results,omitempty" xml:"results>x_hotel,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbhouseAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoXhotelBnbhouseAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbhouseAddAPIResponse)
	},
}

// GetTaobaoXhotelBnbhouseAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbhouseAddAPIResponse
func GetTaobaoXhotelBnbhouseAddAPIResponse() *TaobaoXhotelBnbhouseAddAPIResponse {
	return poolTaobaoXhotelBnbhouseAddAPIResponse.Get().(*TaobaoXhotelBnbhouseAddAPIResponse)
}

// ReleaseTaobaoXhotelBnbhouseAddAPIResponse 将 TaobaoXhotelBnbhouseAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbhouseAddAPIResponse(v *TaobaoXhotelBnbhouseAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbhouseAddAPIResponse.Put(v)
}
