package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品类目属性变更 APIResponse
taobao.item.catprops.modification.get

查询商品类目属性变更信息
*/
type TaobaoItemCatpropsModificationGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemCatpropsModificationGetResponse `json:"taobao_item_catprops_modification_get_response,omitempty"`
}

type TaobaoItemCatpropsModificationGetResponse struct {

    // 返回结果
    Results   []PropsModificationResult `json:"results,omitempty"`

}
