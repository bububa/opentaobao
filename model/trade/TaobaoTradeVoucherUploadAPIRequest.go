package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeVoucherUploadAPIRequest
淘宝交易凭证上传 API请求
taobao.trade.voucher.upload

定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。 */
type TaobaoTradeVoucherUploadAPIRequest struct {
	model.Params
	// 上传文件的名称
	_fileName string
	// 文件
	_fileData *model.File
	// 该笔订单的卖家Nick
	_sellerNick string
	// 该笔订单的买家Nick（混淆nick）
	_buyerNick string
}

// New
