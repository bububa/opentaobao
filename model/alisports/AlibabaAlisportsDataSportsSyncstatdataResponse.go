package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户当天累积数据同步接口 APIResponse
alibaba.alisports.data.sports.syncstatdata

阿里体育数据中心用户当天累积数据同步接口
*/
type AlibabaAlisportsDataSportsSyncstatdataAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDataSportsSyncstatdataResponse
}

type AlibabaAlisportsDataSportsSyncstatdataResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncstatdata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alisp_code
    
    AlispCode   int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`

    
    // alisp_msg
    
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`

    
}
