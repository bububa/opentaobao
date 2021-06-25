package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的创意ID APIResponse
taobao.simba.creativeids.deleted.get

获取删除的创意ID
*/
type TaobaoSimbaCreativeidsDeletedGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCreativeidsDeletedGetResponse `json:"taobao_simba_creativeids_deleted_get_response,omitempty"`
}

type TaobaoSimbaCreativeidsDeletedGetResponse struct {

    // 创意ID列表
    DeletedCreativeIds   []Number `json:"deleted_creative_ids,omitempty"`

}
