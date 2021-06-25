package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询产品线列表 APIResponse
taobao.fenxiao.productcats.get

查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
*/
type TaobaoFenxiaoProductcatsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductcatsGetResponse `json:"taobao_fenxiao_productcats_get_response,omitempty"`
}

type TaobaoFenxiaoProductcatsGetResponse struct {

    // 查询结果记录数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 产品线列表。返回 ProductCat 包含的字段信息。
    Productcats   []ProductCat `json:"productcats,omitempty"`

}
