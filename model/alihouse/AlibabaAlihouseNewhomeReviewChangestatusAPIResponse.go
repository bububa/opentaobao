package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomereviewchangestatusAPIResponse 楼盘测评草稿状态同步 API返回值
// alibaba.alihouse.newhome.review.changestatus
//
// 楼盘测评草稿状态更新
type AlibabaalihousenewhomereviewchangestatusAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomereviewchangestatusAPIResponseModel
}

// AlibabaalihousenewhomereviewchangestatusAPIResponseModel is 楼盘测评草稿状态同步 成功返回结果
type AlibabaalihousenewhomereviewchangestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_review_changestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomereviewchangestatusResult `json:"result,omitempty" xml:"result,omitempty"`
}
