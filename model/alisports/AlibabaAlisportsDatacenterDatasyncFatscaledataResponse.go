package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育接入体脂秤数据 API返回值 
alibaba.alisports.datacenter.datasync.fatscaledata

阿里体育数据中心接入体脂秤数据
*/
type AlibabaAlisportsDatacenterDatasyncFatscaledataAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsDatacenterDatasyncFatscaledataResponse
}

// 阿里体育接入体脂秤数据 成功返回结果
type AlibabaAlisportsDatacenterDatasyncFatscaledataResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_datacenter_datasync_fatscaledata_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回消息
    RsMsg   string `json:"rs_msg,omitempty" xml:"rs_msg,omitempty"`
    // 返回码
    RsCode   string `json:"rs_code,omitempty" xml:"rs_code,omitempty"`
    // 是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 返回对象
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
}
