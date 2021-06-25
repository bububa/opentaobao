package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品导入到渠道 APIRequest
taobao.fenxiao.product.to.channel.import

支持供应商将已有产品导入到某个渠道销售
*/
type TaobaoFenxiaoProductToChannelImportRequest struct {
    model.Params

    // 要导入的渠道[21 零售PLUS]目前仅支持此渠道
    channel   int64 

    // 要导入的产品id
    productId   int64 

}

func NewTaobaoFenxiaoProductToChannelImportRequest() *TaobaoFenxiaoProductToChannelImportRequest{
    return &TaobaoFenxiaoProductToChannelImportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductToChannelImportRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.to.channel.import"
}

func (r TaobaoFenxiaoProductToChannelImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductToChannelImportRequest) SetChannel(channel int64) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoFenxiaoProductToChannelImportRequest) GetChannel() int64 {
    return r.channel
}

func (r *TaobaoFenxiaoProductToChannelImportRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductToChannelImportRequest) GetProductId() int64 {
    return r.productId
}

