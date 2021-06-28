package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增值服务订购服务验证 APIResponse
taobao.vas.service.validate

增值服务订购服务验证
*/
type TaobaoVasServiceValidateAPIResponse struct {
    model.CommonResponse
    TaobaoVasServiceValidateResponse
}

type TaobaoVasServiceValidateResponse struct {
    XMLName xml.Name `xml:"vas_service_validate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // none 没有订购,open 已经开通服务,freeze 服务已经冻结,close 服务已经关闭,error 系统错误
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`

    
}
