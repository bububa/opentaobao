package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosStoreRecordscreenpointinfoAPIRequest
云屏埋点数据记录接口 API请求
alibaba.mos.store.recordscreenpointinfo

记录云屏埋点数据 */
type AlibabaMosStoreRecordscreenpointinfoAPIRequest struct {
	model.Params
	// 云屏埋点信息
	_screenPointInfo string
}

// New
