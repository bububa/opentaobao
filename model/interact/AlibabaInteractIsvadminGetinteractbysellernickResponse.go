package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据sellerNick获取互动实例列表 APIResponse
alibaba.interact.isvadmin.getinteractbysellernick

根据sellerNick获取互动实例列表
*/
type AlibabaInteractIsvadminGetinteractbysellernickAPIResponse struct {
    model.CommonResponse
    AlibabaInteractIsvadminGetinteractbysellernickResponse
}

type AlibabaInteractIsvadminGetinteractbysellernickResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_isvadmin_getinteractbysellernick_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 错误码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`

    
    // 错误信息
    
    Msginfo   string `json:"msginfo,omitempty" xml:"msginfo,omitempty"`

    
    // 返回业务数据
    
    Interactdtos   []InteractDTO `json:"interactdtos,omitempty" xml:"interactdtos>interact_dto,omitempty"`
    
    
}
