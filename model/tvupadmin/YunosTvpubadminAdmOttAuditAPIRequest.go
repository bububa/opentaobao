package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminAdmOttAuditAPIRequest
优酷OTT广告素材审核 API请求
yunos.tvpubadmin.adm.ott.audit

审核优酷OTT端广告素材 */
type YunosTvpubadminAdmOttAuditAPIRequest struct {
	model.Params
	// 广告审核内容，json格式
	_data string
}

// New
