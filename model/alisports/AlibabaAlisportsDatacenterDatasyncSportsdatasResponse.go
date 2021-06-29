package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育运动数据同步统一接口 APIResponse
alibaba.alisports.datacenter.datasync.sportsdatas

给单方提供同步运动数据到阿里体育的接口
*/
type AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDatacenterDatasyncSportsdatasResponse
}

type AlibabaAlisportsDatacenterDatasyncSportsdatasResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_datacenter_datasync_sportsdatas_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回描述
    
    RsMsg   string `json:"rs_msg,omitempty" xml:"rs_msg,omitempty"`

    
    // 返回码
    
    RsCode   string `json:"rs_code,omitempty" xml:"rs_code,omitempty"`

    
    // 是否成功
    
    RsSuccess   bool `json:"rs_success,omitempty" xml:"rs_success,omitempty"`

    
    // 返回值
    
    RsModel   bool `json:"rs_model,omitempty" xml:"rs_model,omitempty"`

    
}
