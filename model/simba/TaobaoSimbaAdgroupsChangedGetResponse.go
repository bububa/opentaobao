package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改的推广组ID和修改时间 APIResponse
taobao.simba.adgroups.changed.get

分页获取修改的推广组ID和修改时间
*/
type TaobaoSimbaAdgroupsChangedGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAdgroupsChangedGetResponse `json:"taobao_simba_adgroups_changed_get_response,omitempty"`
}

type TaobaoSimbaAdgroupsChangedGetResponse struct {

    // 推广组分页对象
    Adgroups   *ADGroupPage `json:"adgroups,omitempty"`

}
