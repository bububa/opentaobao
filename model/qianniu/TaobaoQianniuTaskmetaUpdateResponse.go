package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新任务元数据 APIResponse
taobao.qianniu.taskmeta.update

由任务发起者调用
*/
type TaobaoQianniuTaskmetaUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQianniuTaskmetaUpdateResponse `json:"taobao_qianniu_taskmeta_update_response,omitempty"`
}

type TaobaoQianniuTaskmetaUpdateResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
