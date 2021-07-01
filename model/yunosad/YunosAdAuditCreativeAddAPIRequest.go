package yunosad

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAdAuditCreativeAddAPIRequest
单个创意预审接口 API请求
yunos.ad.audit.creative.add

YunOS广告业务第三方DSP单个创意预审接口 */
type YunosAdAuditCreativeAddAPIRequest struct {
	model.Params
	// 外部dsp的id
	_memberId int64
	// 创意审核入参
	_creative *CreativeParamDto
}

// New
