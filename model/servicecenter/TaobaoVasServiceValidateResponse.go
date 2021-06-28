package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增值服务订购服务验证 APIResponse
taobao.vas.service.validate

增值服务订购服务验证
*/
type TaobaoVasServiceValidateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVasServiceValidateResponse `json:"vas_service_validate_response,omitempty"` 
    TaobaoVasServiceValidateResponse
}

/* model for simplify = false
type TaobaoVasServiceValidateResponse struct {

    // none 没有订购,open 已经开通服务,freeze 服务已经冻结,close 服务已经关闭,error 系统错误
    
    Status   string `json:"status,omitempty"`
    

}
*/

type TaobaoVasServiceValidateResponse struct {

    // none 没有订购,open 已经开通服务,freeze 服务已经冻结,close 服务已经关闭,error 系统错误
    Status   string `json:"status,omitempty"`

}
