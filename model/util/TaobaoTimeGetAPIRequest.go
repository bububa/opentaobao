package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTimeGetAPIRequest
获取淘宝系统当前时间 API请求
taobao.time.get

获取淘宝系统当前时间 */
type TaobaoTimeGetAPIRequest struct {
	model.Params
}

// New
