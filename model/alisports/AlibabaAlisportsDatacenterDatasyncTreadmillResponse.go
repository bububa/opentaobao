package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育同步跑步机设备数据 APIResponse
alibaba.alisports.datacenter.datasync.treadmill

合作方向阿里体育同步跑步机设备的数据
*/
type AlibabaAlisportsDatacenterDatasyncTreadmillAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDatacenterDatasyncTreadmillResponse
}

type AlibabaAlisportsDatacenterDatasyncTreadmillResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_datacenter_datasync_treadmill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回码描述
    
    RsMsg   string `json:"rs_msg,omitempty" xml:"rs_msg,omitempty"`

    
    // 返回码
    
    RsCode   string `json:"rs_code,omitempty" xml:"rs_code,omitempty"`

    
    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 返回值
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
}
