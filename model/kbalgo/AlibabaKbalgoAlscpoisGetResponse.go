package kbalgo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百度批量获取本地poi接口 APIResponse
alibaba.kbalgo.alscpois.get

接口用于百度方获取本地生活poi数据，分页获取。
*/
type AlibabaKbalgoAlscpoisGetAPIResponse struct {
    model.CommonResponse
    AlibabaKbalgoAlscpoisGetResponse
}

type AlibabaKbalgoAlscpoisGetResponse struct {
    XMLName xml.Name `xml:"alibaba_kbalgo_alscpois_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果根节点。
    
    Result   *AlscPoiToBaiduResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
