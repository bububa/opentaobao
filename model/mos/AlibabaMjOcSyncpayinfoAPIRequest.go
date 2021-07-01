package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcSyncpayinfoAPIRequest
支付参考号回传 API请求
alibaba.mj.oc.syncpayinfo

支付参考号同步到oc */
type AlibabaMjOcSyncpayinfoAPIRequest struct {
	model.Params
	// 支付参考号信息
	_posPay *PosPayDto
}

// New
