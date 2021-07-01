package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosStoreGetstorelistAPIRequest
根据屏编号获取专柜集 API请求
alibaba.mos.store.getstorelist

根据屏编号获取专柜集 */
type AlibabaMosStoreGetstorelistAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
}

// New
