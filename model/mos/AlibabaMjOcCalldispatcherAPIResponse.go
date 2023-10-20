package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjoccalldispatcherAPIResponse 呼叫运力 API返回值
// alibaba.mj.oc.calldispatcher
//
// 定时达呼叫运力接口
type AlibabamjoccalldispatcherAPIResponse struct {
	model.CommonResponse
	AlibabamjoccalldispatcherAPIResponseModel
}

// AlibabamjoccalldispatcherAPIResponseModel is 呼叫运力 成功返回结果
type AlibabamjoccalldispatcherAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_calldispatcher_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
