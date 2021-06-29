package ott

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
影视SDK获取设备能力值 APIResponse
youku.ott.alicb.facadeservice.getdata

影视SDK获取设备能力值
*/
type YoukuOttAlicbFacadeserviceGetdataAPIResponse struct {
    model.CommonResponse
    YoukuOttAlicbFacadeserviceGetdataResponse
}

type YoukuOttAlicbFacadeserviceGetdataResponse struct {
    XMLName xml.Name `xml:"youku_ott_alicb_facadeservice_getdata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 设备能力JSON
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
}
