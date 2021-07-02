package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardExpressorderCreateAPIResponse 物流订单创建API API返回值
// tmall.servicecenter.workcard.expressorder.create
//
// 天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
type TmallServicecenterWorkcardExpressorderCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardExpressorderCreateAPIResponseModel
}

// TmallServicecenterWorkcardExpressorderCreateAPIResponseModel is 物流订单创建API 成功返回结果
type TmallServicecenterWorkcardExpressorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_expressorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
