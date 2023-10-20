package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangdeliverystatusupdateAPIResponse 启用/停用配资源 API返回值
// alibaba.dchain.aoxiang.delivery.status.update
//
// 启用/停用配资源
type AlibabadchainaoxiangdeliverystatusupdateAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangdeliverystatusupdateAPIResponseModel
}

// AlibabadchainaoxiangdeliverystatusupdateAPIResponseModel is 启用/停用配资源 成功返回结果
type AlibabadchainaoxiangdeliverystatusupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_delivery_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 启用/停用配资源出参
	DeliveryStatusUpdateResponse *DeliveryStatusUpdateResponse `json:"delivery_status_update_response,omitempty" xml:"delivery_status_update_response,omitempty"`
}
