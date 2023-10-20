package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateAddAPIResponse 新增专享房价 API返回值
// taobao.xhotel.rate.add
//
// 酒店产品库rate添加
type TaobaoXhotelRateAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateAddAPIResponseModel).Reset()
}

// TaobaoXhotelRateAddAPIResponseModel is 新增专享房价 成功返回结果
type TaobaoXhotelRateAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rate_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// results
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// 酒店商品id-酒店rpID
	GidAndRpid string `json:"gid_and_rpid,omitempty" xml:"gid_and_rpid,omitempty"`
	// warnMessage
	WarnMessage string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.GidAndRpid = ""
	m.WarnMessage = ""
}

var poolTaobaoXhotelRateAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateAddAPIResponse)
	},
}

// GetTaobaoXhotelRateAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateAddAPIResponse
func GetTaobaoXhotelRateAddAPIResponse() *TaobaoXhotelRateAddAPIResponse {
	return poolTaobaoXhotelRateAddAPIResponse.Get().(*TaobaoXhotelRateAddAPIResponse)
}

// ReleaseTaobaoXhotelRateAddAPIResponse 将 TaobaoXhotelRateAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateAddAPIResponse(v *TaobaoXhotelRateAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateAddAPIResponse.Put(v)
}
