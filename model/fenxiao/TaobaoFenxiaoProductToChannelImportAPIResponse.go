package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品导入到渠道 API返回值 
taobao.fenxiao.product.to.channel.import

支持供应商将已有产品导入到某个渠道销售
*/
type TaobaoFenxiaoProductToChannelImportAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductToChannelImportAPIResponseModel
}

// 产品导入到渠道 成功返回结果
type TaobaoFenxiaoProductToChannelImportAPIResponseModel struct {
    XMLName xml.Name `xml:"fenxiao_product_to_channel_import_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
