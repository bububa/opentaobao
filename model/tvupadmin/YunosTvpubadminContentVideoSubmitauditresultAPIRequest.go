package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentVideoSubmitauditresultAPIRequest
迎客松提交视频审核结果 API请求
yunos.tvpubadmin.content.video.submitauditresult

迎客松提交视频审核结果 */
type YunosTvpubadminContentVideoSubmitauditresultAPIRequest struct {
	model.Params
	// 审核结果，json格式
	_licenseAuditResult string
}

// New
