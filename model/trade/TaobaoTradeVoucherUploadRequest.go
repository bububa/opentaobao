package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝交易凭证上传 APIRequest
taobao.trade.voucher.upload

定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
*/
type TaobaoTradeVoucherUploadRequest struct {
    model.Params

    // 上传文件的名称
    fileName   string 

    // 文件
    fileData   []*model.File 

    // 该笔订单的卖家Nick
    sellerNick   string 

    // 该笔订单的买家Nick（混淆nick）
    buyerNick   string 

}

func NewTaobaoTradeVoucherUploadRequest() *TaobaoTradeVoucherUploadRequest{
    return &TaobaoTradeVoucherUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeVoucherUploadRequest) GetApiMethodName() string {
    return "taobao.trade.voucher.upload"
}

func (r TaobaoTradeVoucherUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeVoucherUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r TaobaoTradeVoucherUploadRequest) GetFileName() string {
    return r.fileName
}

func (r *TaobaoTradeVoucherUploadRequest) SetFileData(fileData []*model.File) error {
    r.fileData = fileData
    r.Set("file_data", fileData)
    return nil
}

func (r TaobaoTradeVoucherUploadRequest) GetFileData() []*model.File {
    return r.fileData
}

func (r *TaobaoTradeVoucherUploadRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r TaobaoTradeVoucherUploadRequest) GetSellerNick() string {
    return r.sellerNick
}

func (r *TaobaoTradeVoucherUploadRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r TaobaoTradeVoucherUploadRequest) GetBuyerNick() string {
    return r.buyerNick
}

