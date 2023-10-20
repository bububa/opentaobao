package tvpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAppinfoGetAPIResponse tv支付获取应用信息 API返回值
// taobao.tvpay.appinfo.get
//
// tv支付获取应用信息
type TaobaoTvpayAppinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayAppinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTvpayAppinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTvpayAppinfoGetAPIResponseModel).Reset()
}

// TaobaoTvpayAppinfoGetAPIResponseModel is tv支付获取应用信息 成功返回结果
type TaobaoTvpayAppinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_appinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTvpayAppinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTvpayAppinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTvpayAppinfoGetAPIResponse)
	},
}

// GetTaobaoTvpayAppinfoGetAPIResponse 从 sync.Pool 获取 TaobaoTvpayAppinfoGetAPIResponse
func GetTaobaoTvpayAppinfoGetAPIResponse() *TaobaoTvpayAppinfoGetAPIResponse {
	return poolTaobaoTvpayAppinfoGetAPIResponse.Get().(*TaobaoTvpayAppinfoGetAPIResponse)
}

// ReleaseTaobaoTvpayAppinfoGetAPIResponse 将 TaobaoTvpayAppinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTvpayAppinfoGetAPIResponse(v *TaobaoTvpayAppinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoTvpayAppinfoGetAPIResponse.Put(v)
}
