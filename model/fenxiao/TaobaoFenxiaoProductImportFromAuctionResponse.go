package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
导入商品生成产品 APIResponse
taobao.fenxiao.product.import.from.auction

供应商选择关联店铺的前台宝贝，导入生成产品
*/
type TaobaoFenxiaoProductImportFromAuctionAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoFenxiaoProductImportFromAuctionResponse `json:"fenxiao_product_import_from_auction_response,omitempty"` 
    TaobaoFenxiaoProductImportFromAuctionResponse
}

/* model for simplify = false
type TaobaoFenxiaoProductImportFromAuctionResponse struct {

    // 生成的产品id
    
    Pid   int64 `json:"pid,omitempty"`
    

    // 操作时间
    
    OptTime   string `json:"opt_time,omitempty"`
    

}
*/

type TaobaoFenxiaoProductImportFromAuctionResponse struct {

    // 生成的产品id
    Pid   int64 `json:"pid,omitempty"`

    // 操作时间
    OptTime   string `json:"opt_time,omitempty"`

}
