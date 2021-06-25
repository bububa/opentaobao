package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除一个推广组 APIResponse
taobao.simba.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaAdgroupDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAdgroupDeleteResponse `json:"taobao_simba_adgroup_delete_response,omitempty"`
}

type TaobaoSimbaAdgroupDeleteResponse struct {

    // 被删除的推广组
    Adgroup   *ADGroup `json:"adgroup,omitempty"`

}
