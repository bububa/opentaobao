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
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_guoguo_cp_nborderfrontr_updateuser_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    StatusMessage   string `json:"status_message,omitempty" xml:"