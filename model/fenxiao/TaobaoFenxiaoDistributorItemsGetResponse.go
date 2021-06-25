package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品下载记录 APIResponse
taobao.fenxiao.distributor.items.get

供应商查询分销商商品下载记录。
*/
type TaobaoFenxiaoDistributorItemsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoDistributorItemsGetResponse `json:"taobao_fenxiao_distributor_items_get_response,omitempty"`
}

type TaobaoFenxiaoDistributorItemsGetResponse struct {

    // 查询结果记录数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 下载记录对象
    Records   []FenxiaoItemRecord `json:"records,omitempty"`

}
