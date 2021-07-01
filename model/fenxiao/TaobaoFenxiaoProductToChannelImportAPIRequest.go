package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductToChannelImportAPIRequest
产品导入到渠道 API请求
taobao.fenxiao.product.to.channel.import

支持供应商将已有产品导入到某个渠道销售 */
type TaobaoFenxiaoProductToChannelImportAPIRequest struct {
	model.Params
	// 要导入的渠道[21 零售PLUS]目前仅支持此渠道
	_channel int64
	// 要导入的产品id
	_productId int64
}

// New
