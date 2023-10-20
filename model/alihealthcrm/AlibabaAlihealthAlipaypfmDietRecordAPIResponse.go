package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlipaypfmDietRecordAPIResponse 用户每日摄入卡路里总量回传接口 API返回值
// alibaba.alihealth.alipaypfm.diet.record
//
// 用户每日摄入卡路里总量回传接口
type AlibabaAlihealthAlipaypfmDietRecordAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthAlipaypfmDietRecordAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlipaypfmDietRecordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthAlipaypfmDietRecordAPIResponseModel).Reset()
}

// AlibabaAlihealthAlipaypfmDietRecordAPIResponseModel is 用户每日摄入卡路里总量回传接口 成功返回结果
type AlibabaAlihealthAlipaypfmDietRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_diet_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthAlipaypfmDietRecordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthAlipaypfmDietRecordAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthAlipaypfmDietRecordAPIResponse)
	},
}

// GetAlibabaAlihealthAlipaypfmDietRecordAPIResponse 从 sync.Pool 获取 AlibabaAlihealthAlipaypfmDietRecordAPIResponse
func GetAlibabaAlihealthAlipaypfmDietRecordAPIResponse() *AlibabaAlihealthAlipaypfmDietRecordAPIResponse {
	return poolAlibabaAlihealthAlipaypfmDietRecordAPIResponse.Get().(*AlibabaAlihealthAlipaypfmDietRecordAPIResponse)
}

// ReleaseAlibabaAlihealthAlipaypfmDietRecordAPIResponse 将 AlibabaAlihealthAlipaypfmDietRecordAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthAlipaypfmDietRecordAPIResponse(v *AlibabaAlihealthAlipaypfmDietRecordAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthAlipaypfmDietRecordAPIResponse.Put(v)
}
