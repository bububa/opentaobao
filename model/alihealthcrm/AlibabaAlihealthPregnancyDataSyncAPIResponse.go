package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyDataSyncAPIResponse 四类数据同步 API返回值
// alibaba.alihealth.pregnancy.data.sync
//
// 经期调整；基础体温；排卵试纸；B超测排数据同步
type AlibabaAlihealthPregnancyDataSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPregnancyDataSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyDataSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthPregnancyDataSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthPregnancyDataSyncAPIResponseModel is 四类数据同步 成功返回结果
type AlibabaAlihealthPregnancyDataSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_data_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthPregnancyDataSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolAlibabaAlihealthPregnancyDataSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthPregnancyDataSyncAPIResponse)
	},
}

// GetAlibabaAlihealthPregnancyDataSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthPregnancyDataSyncAPIResponse
func GetAlibabaAlihealthPregnancyDataSyncAPIResponse() *AlibabaAlihealthPregnancyDataSyncAPIResponse {
	return poolAlibabaAlihealthPregnancyDataSyncAPIResponse.Get().(*AlibabaAlihealthPregnancyDataSyncAPIResponse)
}

// ReleaseAlibabaAlihealthPregnancyDataSyncAPIResponse 将 AlibabaAlihealthPregnancyDataSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthPregnancyDataSyncAPIResponse(v *AlibabaAlihealthPregnancyDataSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthPregnancyDataSyncAPIResponse.Put(v)
}
