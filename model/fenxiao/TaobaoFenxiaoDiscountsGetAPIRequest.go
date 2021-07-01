package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDiscountsGetAPIRequest
获取折扣信息 API请求
taobao.fenxiao.discounts.get

查询折扣信息 */
type TaobaoFenxiaoDiscountsGetAPIRequest struct {
	model.Params
	// 折扣ID
	_discountId int64
	// 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）
	_extFields string
}

// New
