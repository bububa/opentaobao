package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcOfflineMaxticketnoGetAPIRequest
pos机获取线下最大小票号 API请求
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询 */
type AlibabaMjOcOfflineMaxticketnoGetAPIRequest struct {
	model.Params
	// 收银机号
	_posNo string
	// 外部门店号
	_storeNo string
	// 日期
	_datetime string
}

// New
