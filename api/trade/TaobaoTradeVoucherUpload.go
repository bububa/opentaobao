package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradeVoucherUpload 淘宝交易凭证上传
// taobao.trade.voucher.upload
//
// 定制化交易流程中，涉及到 买家自定义 图片、声音、视频 等多富媒体文件，且该商品或服务的附属sku，通过此接口上传作为交易凭证。
func TaobaoTradeVoucherUpload(clt *core.SDKClient, req *trade.TaobaoTradeVoucherUploadAPIRequest, resp *trade.TaobaoTradeVoucherUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
