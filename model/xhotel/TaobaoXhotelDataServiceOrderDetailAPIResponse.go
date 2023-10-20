package xhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhoteldataserviceorderdetailAPIResponse 服务订单详情 API返回值
// taobao.xhotel.data.service.order.detail
//
// 服务订单详情top接口构建
type TaobaoxhoteldataserviceorderdetailAPIResponse struct {
	model.CommonResponse
	TaobaoxhoteldataserviceorderdetailAPIResponseModel
}

// TaobaoxhoteldataserviceorderdetailAPIResponseModel is 服务订单详情 成功返回结果
type TaobaoxhoteldataserviceorderdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_data_service_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoxhoteldataserviceorderdetailResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
