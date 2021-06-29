package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育接入体脂秤数据 APIResponse
alibaba.alisports.datacenter.datasync.fatscaledata

阿里体育数据中心接入体脂秤数据
*/
type AlibabaAlisportsDatacenterDatasyncFatscaledataAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDatacenterDatasyncFatscaledataResponse
}

type AlibabaAlisportsDatacenterDatasyncFatscaledataResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_datacenter_datasync_fatscaledata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回消息
    
    RsMsg   string `json:"rs_msg,omitempty" xml:"rs_msg,omitempty"`

    
    // 返回码
    
    RsCode   string `json:"rs_code,omitempty" xml:"rs_code,omitempty"`

    
    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 返回对象
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
}
