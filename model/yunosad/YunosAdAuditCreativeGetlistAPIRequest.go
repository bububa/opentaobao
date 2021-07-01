package yunosad

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAdAuditCreativeGetlistAPIRequest
批量获取创意审核状态 API请求
yunos.ad.audit.creative.getlist

批量获取创意审核状态 */
type YunosAdAuditCreativeGetlistAPIRequest struct {
	model.Params
	// 第三方DSP的id
	_memberId int64
	// 状态
	_status string
	// 创意列表
	_creativeIds []string
}

// New
