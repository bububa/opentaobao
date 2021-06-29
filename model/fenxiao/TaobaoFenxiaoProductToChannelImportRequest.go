package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品导入到渠道 API请求
taobao.fenxiao.product.to.channel.import

支持供应商将已有产品导入到某个渠道销售
*/
type TaobaoFenxiaoProductToChannelImportRequest struct {
    model.Params
    // 要导入的渠道[21 零售PLUS]目前仅支持此渠道
    _channel   int64
    // 要导入的产品id
    _productId   int64
}

// 初始化TaobaoFenxiaoProductToChannelImportRequest对象
func NewTaobaoFenxiaoProductToChannelImportRequest() *TaobaoFenxiaoProductToChannelImportRequest{
    return &TaobaoFenxiaoProductToChannelImportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductToChannelImportRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.to.channel.import"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductToChannelImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Channel Setter
// 要导入的渠道[21 零售PLUS]目前仅支持此渠道
func (r *TaobaoFenxiaoProductToChannelImportRequest) SetChannel(_channel int64) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoFenxiaoProductToChannelImportRequest) GetChannel() int64 {
    return r._channel
}
// ProductId Setter
// 要导入的产品id
func (r *TaobaoFenxiaoProductToChannelImportRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductToChannelImportRequest) GetProductId() int64 {
    return r._productId
}
