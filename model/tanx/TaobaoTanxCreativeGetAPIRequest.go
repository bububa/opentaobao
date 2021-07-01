package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxCreativeGetAPIRequest
获取单个审核创意状态 API请求
taobao.tanx.creative.get

获取单个审核创意状态 */
type TaobaoTanxCreativeGetAPIRequest struct {
	model.Params
	// DSP的memberId
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 创意ID
	_creativeId string
}

// New
