package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxAuditDepositcreativeAddAPIRequest
dsp托管创意新增接口 API请求
taobao.tanx.audit.depositcreative.add

dsp托管创意新增接口 */
type TaobaoTanxAuditDepositcreativeAddAPIRequest struct {
	model.Params
	// DSP的memberId
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 托管创意信息
	_creative *CreativeInfoDto
}

// New
