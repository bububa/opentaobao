package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderaftersaleordergetAPIResponse 闲鱼帮卖售后服务单查询 API返回值
// alibaba.idle.tender.aftersale.order.get
//
// 闲鱼帮卖售后服务单查询
type AlibabaidletenderaftersaleordergetAPIResponse struct {
	model.CommonResponse
	AlibabaidletenderaftersaleordergetAPIResponseModel
}

// AlibabaidletenderaftersaleordergetAPIResponseModel is 闲鱼帮卖售后服务单查询 成功返回结果
type AlibabaidletenderaftersaleordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_aftersale_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Module *AlibabaidletenderaftersaleordergetModule `json:"module,omitempty" xml:"module,omitempty"`
	// 查询是否成功
	QuerySuccess bool `json:"query_success,omitempty" xml:"query_success,omitempty"`
}
