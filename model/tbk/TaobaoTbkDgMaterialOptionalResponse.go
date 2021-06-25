package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-物料搜索 APIResponse
taobao.tbk.dg.material.optional

通用物料搜索API（导购）
*/
type TaobaoTbkDgMaterialOptionalAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkDgMaterialOptionalResponse `json:"taobao_tbk_dg_material_optional_response,omitempty"`
}

type TaobaoTbkDgMaterialOptionalResponse struct {

    // 搜索到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // resultList
    ResultList   []MapData `json:"result_list,omitempty"`

    // 本地化-lbs分页标识，请在下一次翻页时作为入参传入
    PageResultKey   string `json:"page_result_key,omitempty"`

}
