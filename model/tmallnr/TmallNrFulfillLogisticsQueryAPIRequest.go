package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillLogisticsQueryAPIRequest
定时送和极速达配送物流信息查询 API请求
tmall.nr.fulfill.logistics.query

发布定时送&极速达物流信息查询服务 */
type TmallNrFulfillLogisticsQueryAPIRequest struct {
	model.Params
	// 交易主订单号
	_mainOrderId int64
	// 业务标识，dss标识定时送业务；jsd表示极速达业务
	_bizIdentity string
}

// New
