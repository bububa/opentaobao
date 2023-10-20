package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse 上门检测服务信息同步 API返回值
// alibaba.alihealth.examination.todoor.serviceinfo.sync
//
// isv同步上门检测服务信息给健康
type AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponseModel is 上门检测服务信息同步 成功返回结果
type AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_todoor_serviceinfo_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse
func GetAlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse() *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse {
	return poolAlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse.Get().(*AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse 将 AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse(v *AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse.Put(v)
}
