package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的推广组ID APIResponse
taobao.simba.adgroupids.deleted.get

获取删除的推广组ID
*/
type TaobaoSimbaAdgroupidsDeletedGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAdgroupidsDeletedGetResponse `json:"taobao_simba_adgroupids_deleted_get_response,omitempty"`
}

type TaobaoSimbaAdgroupidsDeletedGetResponse struct {

    // 推广组ID列表
    DeletedAdgroupIds   []Number `json:"deleted_adgroup_ids,omitempty"`

}
