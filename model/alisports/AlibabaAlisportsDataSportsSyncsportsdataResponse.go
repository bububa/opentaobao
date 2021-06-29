package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户运动数据同步接口 APIResponse
alibaba.alisports.data.sports.syncsportsdata

阿里体育数据中心用户运动数据同步接口
*/
type AlibabaAlisportsDataSportsSyncsportsdataAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDataSportsSyncsportsdataResponse
}

type AlibabaAlisportsDataSportsSyncsportsdataResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncsportsdata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alisp_code
    
    AlispCode   int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`

    
    // alisp_msg
    
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`

    
}
