package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDataSportsSyncsportsdataAPIResponse 阿里体育数据中心用户运动数据同步接口 API返回值
// alibaba.alisports.data.sports.syncsportsdata
//
// 阿里体育数据中心用户运动数据同步接口
type AlibabaAlisportsDataSportsSyncsportsdataAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsDataSportsSyncsportsdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncsportsdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsDataSportsSyncsportsdataAPIResponseModel).Reset()
}

// AlibabaAlisportsDataSportsSyncsportsdataAPIResponseModel is 阿里体育数据中心用户运动数据同步接口 成功返回结果
type AlibabaAlisportsDataSportsSyncsportsdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncsportsdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncsportsdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsDataSportsSyncsportsdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsDataSportsSyncsportsdataAPIResponse)
	},
}

// GetAlibabaAlisportsDataSportsSyncsportsdataAPIResponse 从 sync.Pool 获取 AlibabaAlisportsDataSportsSyncsportsdataAPIResponse
func GetAlibabaAlisportsDataSportsSyncsportsdataAPIResponse() *AlibabaAlisportsDataSportsSyncsportsdataAPIResponse {
	return poolAlibabaAlisportsDataSportsSyncsportsdataAPIResponse.Get().(*AlibabaAlisportsDataSportsSyncsportsdataAPIResponse)
}

// ReleaseAlibabaAlisportsDataSportsSyncsportsdataAPIResponse 将 AlibabaAlisportsDataSportsSyncsportsdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsDataSportsSyncsportsdataAPIResponse(v *AlibabaAlisportsDataSportsSyncsportsdataAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsDataSportsSyncsportsdataAPIResponse.Put(v)
}
