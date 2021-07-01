package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillOrderQueryAPIRequest
零售商获取品牌商的单笔订单 API请求
tmall.nr.fulfill.order.query

零售商获取品牌商的单笔订单，后端服务有零售商和品牌商的绑定关系，存在开关控制；返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位 */
type TmallNrFulfillOrderQueryAPIRequest struct {
	model.Params
	// 业务标识，dss标识定时送业务；jsd表示极速达业务
	_bizIdentity string
	// 交易主订单号
	_orderId int64
	// 预留-扩展信息
	_extParam string
}

// New
