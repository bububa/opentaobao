package yunosad

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAdAuditCreativeGetAPIRequest
获取单个创意审核状态 API请求
yunos.ad.audit.creative.get

获取单个创意审核状态 */
type YunosAdAuditCreativeGetAPIRequest struct {
	model.Params
	// 第三方的dspId
	_memberId int64
	// 第三方广告创意id
	_creativeId string
}

// New
