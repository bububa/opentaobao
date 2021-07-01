package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
导入商品生成产品 API返回值 
taobao.fenxiao.product.import.from.auction

供应商选择关联店铺的前台宝贝，导入生成产品
*/
type TaobaoFenxiaoProductImportFromAuctionAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductImportFromAuctionAPIResponseModel
}

// 导入商品生成产品 成功返回结果
type TaobaoFenxiaoProductImportFromAuctionAPIResponseModel struct {
    XMLName xml.Name `xml:"fenxiao_product_import_from_auction_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 生成的产品id
    Pid   int64 `json:"pid,omitempty" xml:"pid,omitempty"`
    // 操作时间
    OptTime   string `json:"opt_time,omitempty" xml:"opt_time,omitempty"`
}
