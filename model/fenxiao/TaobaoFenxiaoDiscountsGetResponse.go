package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取折扣信息 APIResponse
taobao.fenxiao.discounts.get

查询折扣信息
*/
type TaobaoFenxiaoDiscountsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoDiscountsGetResponse `json:"taobao_fenxiao_discounts_get_response,omitempty"`
}

type TaobaoFenxiaoDiscountsGetResponse struct {

    // 折扣数据结构
    Discounts   []Discount `json:"discounts,omitempty"`

    // 记录数
    TotalResults   int64 `json:"total_results,omitempty"`

}
