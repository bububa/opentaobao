package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
淘宝交易凭证上传 
taobao.trade.voucher.upload

定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
*/
func TaobaoTradeVoucherUpload(clt *core.SDKClient, req *trade.TaobaoTradeVoucherUploadRequest, session string) (*trade.TaobaoTradeVoucherUploadAPIResponse, error) {
    var resp trade.TaobaoTradeVoucherUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
