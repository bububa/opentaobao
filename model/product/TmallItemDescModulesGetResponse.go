package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品描述模块信息获取 APIResponse
tmall.item.desc.modules.get

商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
*/
type TmallItemDescModulesGetAPIResponse struct {
    model.CommonResponse
    Response *TmallItemDescModulesGetResponse `json:"tmall_item_desc_modules_get_response,omitempty"`
}

type TmallItemDescModulesGetResponse struct {

    // 返回描述模块信息
    ModularDescInfo   *ModularDescInfo `json:"modular_desc_info,omitempty"`

}
