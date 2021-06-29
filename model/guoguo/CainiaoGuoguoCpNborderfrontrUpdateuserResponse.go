package guoguo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
小件员信息变更 APIResponse
cainiao.guoguo.cp.nborderfrontr.updateuser

小件员信息变更
*/
type CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse struct {
    model.CommonResponse
    CainiaoGuoguoCpNborderfrontrUpdateuserResponse
}

type CainiaoGuoguoCpNborderfrontrUpdateuserResponse struct {
    XMLName xml.Name `xml:"cainiao_guoguo_cp_nborderfrontr_updateuser_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误信息
    
    StatusMessage   string `json:"status_message,omitempty" xml:"status_message,omitempty"`

    
    // errorCode
    
    StatusCode   string `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
