package ott

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取播放串地址 APIResponse
youku.ott.playservice.getplayurl

获取播放串地址服务
*/
type YoukuOttPlayserviceGetplayurlAPIResponse struct {
    model.CommonResponse
    YoukuOttPlayserviceGetplayurlResponse
}

type YoukuOttPlayserviceGetplayurlResponse struct {
    XMLName xml.Name `xml:"youku_ott_playservice_getplayurl_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *PlayUrlV2Vo `json:"result,omitempty" xml:"result,omitempty"`

    
}
