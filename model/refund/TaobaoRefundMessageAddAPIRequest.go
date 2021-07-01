package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRefundMessageAddAPIRequest
创建退款留言/凭证 API请求
taobao.refund.message.add

创建退款留言/凭证 */
type TaobaoRefundMessageAddAPIRequest struct {
	model.Params
	// 退款编号。
	_refundId int64
	// 留言内容。最大长度: 400个字节
	_content string
	// 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
	_image *model.File
}

// New
