package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotagTagApplyAPIRequest
优惠标签申请 API请求
tmall.promotag.tag.apply

创建优惠标签 */
type TmallPromotagTagApplyAPIRequest struct {
	model.Params
	// 标签名称。注意在UMP中使用新人群标签top变成大写的“NEW_” 如：老标签是top1234，新标签是NEW_1234 。
	_tagName string
	// 标签用途描述
	_tagDesc string
	// 标签开始时间
	_startTime string
	// 标签结束时间
	_endTime string
}

// New
