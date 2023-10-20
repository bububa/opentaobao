package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkelemebilldetailgetAPIResponse 饿了么对账单查询，带订单明细 API返回值
// alibaba.wdk.eleme.bill.detail.get
//
// 查询饿了么对账单信息，带订单明细
type AlibabawdkelemebilldetailgetAPIResponse struct {
	model.CommonResponse
	AlibabawdkelemebilldetailgetAPIResponseModel
}

// AlibabawdkelemebilldetailgetAPIResponseModel is 饿了么对账单查询，带订单明细 成功返回结果
type AlibabawdkelemebilldetailgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_eleme_bill_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabawdkelemebilldetailgetApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
