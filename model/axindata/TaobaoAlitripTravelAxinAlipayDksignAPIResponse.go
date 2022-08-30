package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinAlipayDksignAPIResponse 支付宝代扣签约 API返回值
// taobao.alitrip.travel.axin.alipay.dksign
//
// 提供支付宝代扣签约服务
type TaobaoAlitripTravelAxinAlipayDksignAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelAxinAlipayDksignAPIResponseModel
}

// TaobaoAlitripTravelAxinAlipayDksignAPIResponseModel is 支付宝代扣签约 成功返回结果
type TaobaoAlitripTravelAxinAlipayDksignAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_axin_alipay_dksign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripTravelAxinAlipayDksignResult `json:"result,omitempty" xml:"result,omitempty"`
}
