package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDataSportsSyncuserdataAPIResponse 阿里体育数据中心用户个人信息同步接口 API返回值
// alibaba.alisports.data.sports.syncuserdata
//
// 阿里体育数据中心用户个人信息同步接口
type AlibabaAlisportsDataSportsSyncuserdataAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsDataSportsSyncuserdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncuserdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsDataSportsSyncuserdataAPIResponseModel).Reset()
}

// AlibabaAlisportsDataSportsSyncuserdataAPIResponseModel is 阿里体育数据中心用户个人信息同步接口 成功返回结果
type AlibabaAlisportsDataSportsSyncuserdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncuserdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsDataSportsSyncuserdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsDataSportsSyncuserdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsDataSportsSyncuserdataAPIResponse)
	},
}

// GetAlibabaAlisportsDataSportsSyncuserdataAPIResponse 从 sync.Pool 获取 AlibabaAlisportsDataSportsSyncuserdataAPIResponse
func GetAlibabaAlisportsDataSportsSyncuserdataAPIResponse() *AlibabaAlisportsDataSportsSyncuserdataAPIResponse {
	return poolAlibabaAlisportsDataSportsSyncuserdataAPIResponse.Get().(*AlibabaAlisportsDataSportsSyncuserdataAPIResponse)
}

// ReleaseAlibabaAlisportsDataSportsSyncuserdataAPIResponse 将 AlibabaAlisportsDataSportsSyncuserdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsDataSportsSyncuserdataAPIResponse(v *AlibabaAlisportsDataSportsSyncuserdataAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsDataSportsSyncuserdataAPIResponse.Put(v)
}
