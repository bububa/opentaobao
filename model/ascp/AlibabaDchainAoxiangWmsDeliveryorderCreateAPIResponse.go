package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangwmsdeliveryordercreateAPIResponse 回传仓库接单通知 API返回值
// alibaba.dchain.aoxiang.wms.deliveryorder.create
//
// WMS上报仓库接单节点状态信息，代表接单环节。
type AlibabadchainaoxiangwmsdeliveryordercreateAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangwmsdeliveryordercreateAPIResponseModel
}

// AlibabadchainaoxiangwmsdeliveryordercreateAPIResponseModel is 回传仓库接单通知 成功返回结果
type AlibabadchainaoxiangwmsdeliveryordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_wms_deliveryorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传结果
	DeliveryOrderReportResponse *DeliveryOrderReportResponse `json:"delivery_order_report_response,omitempty" xml:"delivery_order_report_response,omitempty"`
}
