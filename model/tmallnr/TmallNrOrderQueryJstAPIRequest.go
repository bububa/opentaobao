package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrOrderQueryJstAPIRequest
获取同城配送业务单笔订单 API请求
tmall.nr.order.query.jst

同城配送业务获取单笔订单 */
type TmallNrOrderQueryJstAPIRequest struct {
	model.Params
	// 业务标识，dss标识定时送业务；jsd表示极速达业务
	_bizIdentity string
	// 交易主订单号
	_orderId int64
	// 预留-扩展信息
	_extParam string
}

// New
