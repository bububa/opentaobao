package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取修改的推广组ID APIResponse
taobao.simba.adgroupids.changed.get

获取修改的推广组ID
*/
type TaobaoSimbaAdgroupidsChangedGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAdgroupidsChangedGetResponse `json:"taobao_simba_adgroupids_changed_get_response,omitempty"`
}

type TaobaoSimbaAdgroupidsChangedGetResponse struct {

    // 推广组ID列表
    ChangedAdgroupids   []Number `json:"changed_adgroupids,omitempty"`

}
