package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU商家编码更新接口 APIResponse
tmall.item.outerid.update

天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
*/
type TmallItemOuteridUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallItemOuteridUpdateResponse `json:"tmall_item_outerid_update_response,omitempty"`
}

type TmallItemOuteridUpdateResponse struct {

    // 商家编码更新结果
    OuteridUpdateResult   string `json:"outerid_update_result,omitempty"`

}
