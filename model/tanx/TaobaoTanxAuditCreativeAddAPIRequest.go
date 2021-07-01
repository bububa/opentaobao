package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxAuditCreativeAddAPIRequest
创意预审新增接口 API请求
taobao.tanx.audit.creative.add

创意预审新增接口 */
type TaobaoTanxAuditCreativeAddAPIRequest struct {
	model.Params
	// DSP的memberId
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 预审核创意对象
	_creative *CreativeParamDto
}

// New
