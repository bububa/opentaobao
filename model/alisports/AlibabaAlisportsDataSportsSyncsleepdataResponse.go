package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户睡眠数据同步接口 API返回值 
alibaba.alisports.data.sports.syncsleepdata

阿里体育数据中心用户睡眠数据同步接口
*/
type AlibabaAlisportsDataSportsSyncsleepdataAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDataSportsSyncsleepdataResponse
}

// 阿里体育数据中心用户睡眠数据同步接口 成功返回结果
type AlibabaAlisportsDataSportsSyncsleepdataResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncsleepdata_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alisp_code
    AlispCode   int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
    // alisp_msg
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
}
