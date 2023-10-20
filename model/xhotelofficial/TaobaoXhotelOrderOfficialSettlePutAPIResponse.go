package xhotelofficial

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfficialSettlePutAPIResponse 官网信用住结账接口 API返回值
// taobao.xhotel.order.official.settle.put
//
// 用于酒店官网信用住商家结账调用
type TaobaoXhotelOrderOfficialSettlePutAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderOfficialSettlePutAPIResponseModel
}

// TaobaoXhotelOrderOfficialSettlePutAPIResponseModel is 官网信用住结账接口 成功返回结果
type TaobaoXhotelOrderOfficialSettlePutAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_official_settle_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
