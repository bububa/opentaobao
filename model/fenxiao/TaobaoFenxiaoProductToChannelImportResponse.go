package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品导入到渠道 APIResponse
taobao.fenxiao.product.to.channel.import

支持供应商将已有产品导入到某个渠道销售
*/
type TaobaoFenxiaoProductToChannelImportAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductToChannelImportResponse
}

type TaobaoFenxiaoProductToChannelImportResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_to_channel_import_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

}
