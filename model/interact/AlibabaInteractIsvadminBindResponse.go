package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建及绑定互动实例 APIResponse
alibaba.interact.isvadmin.bind

创建互动实例，并绑定奖池
*/
type AlibabaInteractIsvadminBindAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractIsvadminBindResponse `json:"alibaba_interact_isvadmin_bind_response,omitempty"`
}

type AlibabaInteractIsvadminBindResponse struct {

    // 返回创建并且绑定成功的互动实例
    Result   *CommonResult `json:"result,omitempty"`

}
