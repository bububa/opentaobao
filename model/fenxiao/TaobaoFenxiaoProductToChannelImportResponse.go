package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品导入到渠道 APIResponse
taobao.fenxiao.product.to.channel.import

支持供应商将已有产品导入到某个渠道销售
*/
type TaobaoFenxiaoProductToChannelImportAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductToChannelImportResponse `json:"taobao_fenxiao_product_to_channel_import_response,omitempty"`
}

type TaobaoFenxiaoProductToChannelImportResponse struct {

}
