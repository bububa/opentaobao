package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝交易凭证上传 API请求
taobao.trade.voucher.upload

定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
*/
type TaobaoTradeVoucherUploadAPIRequest struct {
    model.Params
    // 上传文件的名称
    _fileName   string
    // 文件
    _fileData   *model.File
    // 该笔订单的卖家Nick
    _sellerNick   string
    // 该笔订单的买家Nick（混淆nick）
    _buyerNick   string
}

// 初始化TaobaoTradeVoucherUploadAPIRequest对象
func NewTaobaoTradeVoucherUploadRequest() *TaobaoTradeVoucherUploadAPIRequest{
    return &TaobaoTradeVoucherUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeVoucherUploadAPIRequest) GetApiMethodName() string {
    return "taobao.trade.voucher.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeVoucherUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileName Setter
// 上传文件的名称
func (r *TaobaoTradeVoucherUploadAPIRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetFileName() string {
    return r._fileName
}
// FileData Setter
// 文件
func (r *TaobaoTradeVoucherUploadAPIRequest) SetFileData(_fileData *model.File) error {
    r._fileData = _fileData
    r.Set("file_data", _fileData)
    return nil
}

// FileData Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetFileData() *model.File {
    return r._fileData
}
// SellerNick Setter
// 该笔订单的卖家Nick
func (r *TaobaoTradeVoucherUploadAPIRequest) SetSellerNick(_sellerNick string) error {
    r._sellerNick = _sellerNick
    r.Set("seller_nick", _sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetSellerNick() string {
    return r._sellerNick
}
// BuyerNick Setter
// 该笔订单的买家Nick（混淆nick）
func (r *TaobaoTradeVoucherUploadAPIRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetBuyerNick() string {
    return r._buyerNick
}
