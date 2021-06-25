package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据sellerNick获取互动实例列表 APIResponse
alibaba.interact.isvadmin.getinteractbysellernick

根据sellerNick获取互动实例列表
*/
type AlibabaInteractIsvadminGetinteractbysellernickAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractIsvadminGetinteractbysellernickResponse `json:"alibaba_interact_isvadmin_getinteractbysellernick_response,omitempty"`
}

type AlibabaInteractIsvadminGetinteractbysellernickResponse struct {

    // 返回结果是否成功
    Succ   bool `json:"succ,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

    // 错误信息
    Msginfo   string `json:"msginfo,omitempty"`

    // 返回业务数据
    Interactdtos   []InteractDTO `json:"interactdtos,omitempty"`

}
