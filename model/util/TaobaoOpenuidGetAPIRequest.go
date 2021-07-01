package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenuidGetAPIRequest
获取授权账号对应的OpenUid API请求
taobao.openuid.get

获取授权账号对应的OpenUid */
type TaobaoOpenuidGetAPIRequest struct {
	model.Params
}

// New
