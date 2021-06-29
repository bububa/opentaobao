package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建及绑定互动实例 APIResponse
alibaba.interact.isvadmin.bind

创建互动实例，并绑定奖池
*/
type AlibabaInteractIsvadminBindAPIResponse struct {
    model.CommonResponse
    AlibabaInteractIsvadminBindResponse
}

type AlibabaInteractIsvadminBindResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_isvadmin_bind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回创建并且绑定成功的互动实例
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
