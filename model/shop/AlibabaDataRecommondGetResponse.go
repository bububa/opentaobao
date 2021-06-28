package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推荐信息 APIResponse
alibaba.data.recommond.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaDataRecommondGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_data_recommond_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 客户端鉴权虚拟api使用
    
    Unnamed   string `json:"unnamed,omitempty" xml:"