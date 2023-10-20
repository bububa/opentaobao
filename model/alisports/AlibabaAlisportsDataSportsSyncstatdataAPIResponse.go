package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDataSportsSyncstatdataAPIResponse 阿里体育数据中心用户当天累积数据同步接口 API返回值
// alibaba.alisports.data.sports.syncstatdata
//
// 阿里体育数据中心用户当天累积数据同步接口
type AlibabaAlisportsDataSportsSyncstatdataAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsDataSportsSyncstatdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncstatdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsDataSportsSyncstatdataAPIResponseModel).Reset()
}

// AlibabaAlisportsDataSportsSyncstatdataAPIResponseModel is 阿里体育数据中心用户当天累积数据同步接口 成功返回结果
type AlibabaAlisportsDataSportsSyncstatdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncstatdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncstatdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsDataSportsSyncstatdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsDataSportsSyncstatdataAPIResponse)
	},
}

// GetAlibabaAlisportsDataSportsSyncstatdataAPIResponse 从 sync.Pool 获取 AlibabaAlisportsDataSportsSyncstatdataAPIResponse
func GetAlibabaAlisportsDataSportsSyncstatdataAPIResponse() *AlibabaAlisportsDataSportsSyncstatdataAPIResponse {
	return poolAlibabaAlisportsDataSportsSyncstatdataAPIResponse.Get().(*AlibabaAlisportsDataSportsSyncstatdataAPIResponse)
}

// ReleaseAlibabaAlisportsDataSportsSyncstatdataAPIResponse 将 AlibabaAlisportsDataSportsSyncstatdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsDataSportsSyncstatdataAPIResponse(v *AlibabaAlisportsDataSportsSyncstatdataAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsDataSportsSyncstatdataAPIResponse.Put(v)
}
