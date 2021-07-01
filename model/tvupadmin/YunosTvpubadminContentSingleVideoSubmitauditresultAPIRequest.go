package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest
单视频审核提交审核结果 API请求
yunos.tvpubadmin.content.single.video.submitauditresult

单视频审核提交审核结果 */
type YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest struct {
	model.Params
	// 审核人
	_auditor string
	// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
	_licenseState int64
	// 牌照方
	_license int64
	// 备注说明
	_auditComment string
	// 视频审核ID
	_videoAuditId int64
}

// New
