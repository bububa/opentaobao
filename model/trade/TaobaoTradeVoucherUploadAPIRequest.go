package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradevoucheruploadAPIRequest 淘宝交易凭证上传 API请求
// taobao.trade.voucher.upload
//
// 定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
type TaobaotradevoucheruploadAPIRequest struct {
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

// NewTaobaotradevoucheruploadRequest 初始化TaobaotradevoucheruploadAPIRequest对象
func NewTaobaotradevoucheruploadRequest() *TaobaotradevoucheruploadAPIRequest {
	return &TaobaotradevoucheruploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradevoucheruploadAPIRequest) GetApiMethodName() string {
	return "taobao.trade.voucher.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradevoucheruploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradevoucheruploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 上传文件的名称
func (r *TaobaotradevoucheruploadAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r TaobaotradevoucheruploadAPIRequest) GetFileName() string {
	return r._fileName
}

// SetSellerNick is SellerNick Setter
// 该笔订单的卖家Nick
func (r *TaobaotradevoucheruploadAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaotradevoucheruploadAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetBuyerNick is BuyerNick Setter
// 该笔订单的买家Nick（混淆nick）
func (r *TaobaotradevoucheruploadAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TaobaotradevoucheruploadAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetFileData is FileData Setter
// 文件
func (r *TaobaotradevoucheruploadAPIRequest) SetFileData(_fileData *model.File) error {
	r._fileData = _fileData
	r.Set("file_data", _fileData)
	return nil
}

// GetFileData FileData Getter
func (r TaobaotradevoucheruploadAPIRequest) GetFileData() *model.File {
	return r._fileData
}
