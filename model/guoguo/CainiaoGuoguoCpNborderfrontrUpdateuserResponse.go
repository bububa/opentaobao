package guoguo

import (
    "github.com/bububa/opentaobao/model"
)

/* 
小件员信息变更 APIResponse
cainiao.guoguo.cp.nborderfrontr.updateuser

小件员信息变更
*/
type CainiaoGuoguoCpNborderfrontrUpdateuserAPIResponse struct {
    model.CommonResponse
    Response *CainiaoGuoguoCpNborderfrontrUpdateuserResponse `json:"cainiao_guoguo_cp_nborderfrontr_updateuser_response,omitempty"`
}

type CainiaoGuoguoCpNborderfrontrUpdateuserResponse struct {

    // 错误信息
    StatusMessage   string `json:"status_message,omitempty"`

    // errorCode
    StatusCode   string `json:"status_code,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
