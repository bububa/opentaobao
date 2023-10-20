package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportsdatasportssyncstatdataAPIResponse 阿里体育数据中心用户当天累积数据同步接口 API返回值
// alibaba.alisports.data.sports.syncstatdata
//
// 阿里体育数据中心用户当天累积数据同步接口
type AlibabaalisportsdatasportssyncstatdataAPIResponse struct {
	model.CommonResponse
	AlibabaalisportsdatasportssyncstatdataAPIResponseModel
}

// AlibabaalisportsdatasportssyncstatdataAPIResponseModel is 阿里体育数据中心用户当天累积数据同步接口 成功返回结果
type AlibabaalisportsdatasportssyncstatdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncstatdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}
