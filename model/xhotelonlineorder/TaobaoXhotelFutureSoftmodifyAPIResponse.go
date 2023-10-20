package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelFutureSoftmodifyAPIResponse 未来酒店信息下发 API返回值
// taobao.xhotel.future.softmodify
//
// 未来酒店信息下发，包含PMS订单查询和自助入住
type TaobaoXhotelFutureSoftmodifyAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelFutureSoftmodifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelFutureSoftmodifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelFutureSoftmodifyAPIResponseModel).Reset()
}

// TaobaoXhotelFutureSoftmodifyAPIResponseModel is 未来酒店信息下发 成功返回结果
type TaobaoXhotelFutureSoftmodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_future_softmodify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求报文示例：https://open.alitrip.com/docs/doc.htm?spm=a21tt.7629140.0.0.Q8jazn&amp;docType=1&amp;articleId=104398中的自助入住请求示例(升级版)和酒店PMS信息查询 响应报文示例：https://open.alitrip.com/docs/doc.htm?docType=1&amp;articleId=106153中的自助checkIn订单信息上传和通用PMS结果查询结果
	Context string `json:"context,omitempty" xml:"context,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelFutureSoftmodifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Context = ""
}

var poolTaobaoXhotelFutureSoftmodifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelFutureSoftmodifyAPIResponse)
	},
}

// GetTaobaoXhotelFutureSoftmodifyAPIResponse 从 sync.Pool 获取 TaobaoXhotelFutureSoftmodifyAPIResponse
func GetTaobaoXhotelFutureSoftmodifyAPIResponse() *TaobaoXhotelFutureSoftmodifyAPIResponse {
	return poolTaobaoXhotelFutureSoftmodifyAPIResponse.Get().(*TaobaoXhotelFutureSoftmodifyAPIResponse)
}

// ReleaseTaobaoXhotelFutureSoftmodifyAPIResponse 将 TaobaoXhotelFutureSoftmodifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelFutureSoftmodifyAPIResponse(v *TaobaoXhotelFutureSoftmodifyAPIResponse) {
	v.Reset()
	poolTaobaoXhotelFutureSoftmodifyAPIResponse.Put(v)
}
