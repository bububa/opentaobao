package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxAuditCreativeModifyAPIRequest
创意修改接口 API请求
taobao.tanx.audit.creative.modify

创意修改接口 */
type TaobaoTanxAuditCreativeModifyAPIRequest struct {
	model.Params
	// DSP用户ID
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
}

// New
