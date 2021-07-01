package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户比赛数据同步接口 API返回值 
alibaba.alisports.data.sports.syncmatchdata

阿里体育数据中心用户比赛数据同步接口
*/
type AlibabaAlisportsDataSportsSyncmatchdataAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDataSportsSyncmatchdataAPIResponseModel
}

// 阿里体育数据中心用户比赛数据同步接口 成功返回结果
type AlibabaAlisportsDataSportsSyncmatchdataAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alisports_data_sports_syncmatchdata_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应码
    AlispCode   int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
    // 响应信息
    AlispMsg   string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
}
