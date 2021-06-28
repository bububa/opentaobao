package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取可用宝贝描述规范化模块 APIResponse
taobao.item.anchor.get

根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
*/
type TaobaoItemAnchorGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemAnchorGetResponse `json:"item_anchor_get_response,omitempty"` 
    TaobaoItemAnchorGetResponse
}

/* model for simplify = false
type TaobaoItemAnchorGetResponse struct {

    // 返回的宝贝描述模板结果数目
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 宝贝描述规范化可使用打标模块的锚点信息
    
    AnchorModules  struct {
        IdsModule  []IdsModule `json:"ids_module,omitempty"`
    } `json:"anchor_modules,omitempty"`
    

}
*/

type TaobaoItemAnchorGetResponse struct {

    // 返回的宝贝描述模板结果数目
    TotalResults   int64 `json:"total_results,omitempty"`

    // 宝贝描述规范化可使用打标模块的锚点信息
    AnchorModules   []IdsModule `json:"anchor_modules,omitempty"`

}
