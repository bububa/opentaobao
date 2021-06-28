package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取DSP用户的创意审核结果 APIResponse
taobao.tanx.creatives.get

批量获取DSP用户的创意审核结果
*/
type TaobaoTanxCreativesGetAPIResponse struct {
    model.CommonResponse
    TaobaoTanxCreativesGetResponse
}

type TaobaoTanxCreativesGetResponse struct {
    XMLName xml.Name `xml:"tanx_creatives_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用的成功信息或失败信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 调用返回码
    
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty" xml:"tanx_error_code,omitempty"`

    
    // 调用是否成功
    
    IsOk   bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`

    
    // 返回的创意列表
    
    Results   []CreativeDto `json:"results,omitempty" xml:"results>creative_dto,omitempty"`
    
    
}
