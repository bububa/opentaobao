package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectlotteryresultsubmitAPIResponse 楼盘摇号结果提交 API返回值
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
type AlibabaalihousenewhomeprojectlotteryresultsubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectlotteryresultsubmitAPIResponseModel
}

// AlibabaalihousenewhomeprojectlotteryresultsubmitAPIResponseModel is 楼盘摇号结果提交 成功返回结果
type AlibabaalihousenewhomeprojectlotteryresultsubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_lottery_result_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaalihousenewhomeprojectlotteryresultsubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
