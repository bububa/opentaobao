package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosStoreGetcloudshelfversionAPIRequest
获取云货架版本信息 API请求
alibaba.mos.store.getcloudshelfversion

根据屏编号获取云货架版本信息 */
type AlibabaMosStoreGetcloudshelfversionAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
}

// New
