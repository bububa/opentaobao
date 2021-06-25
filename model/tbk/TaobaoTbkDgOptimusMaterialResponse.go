package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-物料精选 APIResponse
taobao.tbk.dg.optimus.material

支持入参对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
*/
type TaobaoTbkDgOptimusMaterialAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTbkDgOptimusMaterialResponse `json:"taobao_tbk_dg_optimus_material_response,omitempty"`
}

type TaobaoTbkDgOptimusMaterialResponse struct {

    // resultList
    ResultList   []MapData `json:"result_list,omitempty"`

    // 推荐信息-是否抄底
    IsDefault   string `json:"is_default,omitempty"`

    // 商品总数-目前只有全品库商品查询有该字段
    TotalCount   int64 `json:"total_count,omitempty"`

}
