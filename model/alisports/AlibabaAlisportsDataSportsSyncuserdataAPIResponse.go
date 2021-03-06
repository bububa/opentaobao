package alisports

import (
	"encoding/xml"

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
