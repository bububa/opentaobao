package youkudsp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷实时批量获取可投放设备资源 APIResponse
youku.dsp.delivery.resource.multiget

优酷实时获取可投放设备资源信息,为第三方渠道提供素材获取人群识别的api,支持批量获取
*/
type YoukuDspDeliveryResourceMultigetAPIResponse struct {
    model.CommonResponse
    YoukuDspDeliveryResourceMultigetResponse
}

type YoukuDspDeliveryResourceMultigetResponse struct {
    XMLName xml.Name `xml:"youku_dsp_delivery_resource_multiget_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 素材列表
    
    Models   []DeliveryList `json:"models,omitempty" xml:"models>delivery_list,omitempty"`
    
    
    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 额外信息
    
    ExtraInfo   string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`

    
    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 是否成功
    
    SuccessFlag   bool `json:"success_flag,omitempty" xml:"success_flag,omitempty"`

    
}
