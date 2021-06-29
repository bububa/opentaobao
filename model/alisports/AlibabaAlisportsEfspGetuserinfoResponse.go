package alisports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户详细信息 APIResponse
alibaba.alisports.efsp.getuserinfo

阿里体育-体育健身-获取用户详细信息
*/
type AlibabaAlisportsEfspGetuserinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlisportsEfspGetuserinfoResponse
}

type AlibabaAlisportsEfspGetuserinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alisports_efsp_getuserinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TrilateralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
