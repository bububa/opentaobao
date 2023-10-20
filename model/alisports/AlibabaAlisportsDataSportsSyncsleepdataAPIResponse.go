package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDataSportsSyncsleepdataAPIResponse 阿里体育数据中心用户睡眠数据同步接口 API返回值
// alibaba.alisports.data.sports.syncsleepdata
//
// 阿里体育数据中心用户睡眠数据同步接口
type AlibabaAlisportsDataSportsSyncsleepdataAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsDataSportsSyncsleepdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncsleepdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsDataSportsSyncsleepdataAPIResponseModel).Reset()
}

// AlibabaAlisportsDataSportsSyncsleepdataAPIResponseModel is 阿里体育数据中心用户睡眠数据同步接口 成功返回结果
type AlibabaAlisportsDataSportsSyncsleepdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncsleepdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncsleepdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsDataSportsSyncsleepdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsDataSportsSyncsleepdataAPIResponse)
	},
}

// GetAlibabaAlisportsDataSportsSyncsleepdataAPIResponse 从 sync.Pool 获取 AlibabaAlisportsDataSportsSyncsleepdataAPIResponse
func GetAlibabaAlisportsDataSportsSyncsleepdataAPIResponse() *AlibabaAlisportsDataSportsSyncsleepdataAPIResponse {
	return poolAlibabaAlisportsDataSportsSyncsleepdataAPIResponse.Get().(*AlibabaAlisportsDataSportsSyncsleepdataAPIResponse)
}

// ReleaseAlibabaAlisportsDataSportsSyncsleepdataAPIResponse 将 AlibabaAlisportsDataSportsSyncsleepdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsDataSportsSyncsleepdataAPIResponse(v *AlibabaAlisportsDataSportsSyncsleepdataAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsDataSportsSyncsleepdataAPIResponse.Put(v)
}
