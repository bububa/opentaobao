package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeVoucherUploadAPIRequest 淘宝交易凭证上传 API请求
// taobao.trade.voucher.upload
//
// 定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
type TaobaoTradeVoucherUploadAPIRequest struct {
	model.Params
	// 上传文件的名称
	_fileName string
	// 该笔订单的卖家Nick
	_sellerNick string
	// 该笔订单的买家Nick（混淆nick）
	_buyerNick string
	// 文件
	_fileData *model.File
}

// NewTaobaoTradeVoucherUploadRequest 初始化TaobaoTradeVoucherUploadAPIRequest对象
func NewTaobaoTradeVoucherUploadRequest() *TaobaoTradeVoucherUploadAPIRequest {
	return &TaobaoTradeVoucherUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeVoucherUploadAPIRequest) Reset() {
	r._fileName = ""
	r._sellerNick = ""
	r._buyerNick = ""
	r._fileData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeVoucherUploadAPIRequest) GetApiMethodName() string {
	return "taobao.trade.voucher.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeVoucherUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeVoucherUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 上传文件的名称
func (r *TaobaoTradeVoucherUploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetSellerNick is SellerNick Setter
// 该笔订单的卖家Nick
func (r *TaobaoTradeVoucherUploadAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetBuyerNick is BuyerNick Setter
// 该笔订单的买家Nick（混淆nick）
func (r *TaobaoTradeVoucherUploadAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetFileData is FileData Setter
// 文件
func (r *TaobaoTradeVoucherUploadAPIRequest) SetFileData(_fileData *model.File) error {
	r._fileData = _fileData
	r.Set("file_data", _fileData)
	return nil
}

// GetFileData FileData Getter
func (r TaobaoTradeVoucherUploadAPIRequest) GetFileData() *model.File {
	return r._fileData
}

var poolTaobaoTradeVoucherUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeVoucherUploadRequest()
	},
}

// GetTaobaoTradeVoucherUploadRequest 从 sync.Pool 获取 TaobaoTradeVoucherUploadAPIRequest
func GetTaobaoTradeVoucherUploadAPIRequest() *TaobaoTradeVoucherUploadAPIRequest {
	return poolTaobaoTradeVoucherUploadAPIRequest.Get().(*TaobaoTradeVoucherUploadAPIRequest)
}

// ReleaseTaobaoTradeVoucherUploadAPIRequest 将 TaobaoTradeVoucherUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeVoucherUploadAPIRequest(v *TaobaoTradeVoucherUploadAPIRequest) {
	v.Reset()
	poolTaobaoTradeVoucherUploadAPIRequest.Put(v)
}
