package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopOpenidConvertAPIRequest
混淆nick转openid API请求
taobao.top.openid.convert

混淆nick转openid，生成混淆nick必须与当前请求的isv匹配 */
type TaobaoTopOpenidConvertAPIRequest struct {
	model.Params
	// 混淆nick转open_id
	_mixNick string
}

// New
