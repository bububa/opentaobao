package tvpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAccessDataGetAPIResponse tv支付 API返回值
// taobao.tvpay.access.data.get
//
// 在匿名用户支付后尝试为其登陆绑定的淘宝账号
type TaobaoTvpayAccessDataGetAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayAccessDataGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTvpayAccessDataGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTvpayAccessDataGetAPIResponseModel).Reset()
}

// TaobaoTvpayAccessDataGetAPIResponseModel is tv支付 成功返回结果
type TaobaoTvpayAccessDataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_access_data_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTvpayAccessDataGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTvpayAccessDataGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTvpayAccessDataGetAPIResponse)
	},
}

// GetTaobaoTvpayAccessDataGetAPIResponse 从 sync.Pool 获取 TaobaoTvpayAccessDataGetAPIResponse
func GetTaobaoTvpayAccessDataGetAPIResponse() *TaobaoTvpayAccessDataGetAPIResponse {
	return poolTaobaoTvpayAccessDataGetAPIResponse.Get().(*TaobaoTvpayAccessDataGetAPIResponse)
}

// ReleaseTaobaoTvpayAccessDataGetAPIResponse 将 TaobaoTvpayAccessDataGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTvpayAccessDataGetAPIResponse(v *TaobaoTvpayAccessDataGetAPIResponse) {
	v.Reset()
	poolTaobaoTvpayAccessDataGetAPIResponse.Put(v)
}
