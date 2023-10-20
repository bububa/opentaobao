package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse 楼盘摇号结果提交 API返回值
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
type AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponseModel is 楼盘摇号结果提交 成功返回结果
type AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_lottery_result_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	Result *AlibabaAlihouseNewhomeProjectLotteryResultSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse
func GetAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse() *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse 将 AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse(v *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIResponse.Put(v)
}
