package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressModifyAppointAPIResponse 快递改约api API返回值
// taobao.logistics.express.modify.appoint
//
// 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
type TaobaoLogisticsExpressModifyAppointAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressModifyAppointAPIResponseModel
}

// TaobaoLogisticsExpressModifyAppointAPIResponseModel is 快递改约api 成功返回结果
type TaobaoLogisticsExpressModifyAppointAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_modify_appoint_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *SingleResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
