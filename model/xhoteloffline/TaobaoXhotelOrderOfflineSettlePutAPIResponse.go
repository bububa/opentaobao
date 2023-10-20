package xhoteloffline

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderOfflineSettlePutAPIResponse 线下信用住结账专用接口 API返回值
// taobao.xhotel.order.offline.settle.put
//
// 线下信用住结账专用接口
type TaobaoXhotelOrderOfflineSettlePutAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderOfflineSettlePutAPIResponseModel
}

// TaobaoXhotelOrderOfflineSettlePutAPIResponseModel is 线下信用住结账专用接口 成功返回结果
type TaobaoXhotelOrderOfflineSettlePutAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_offline_settle_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
