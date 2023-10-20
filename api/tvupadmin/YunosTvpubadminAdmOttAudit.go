package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminAdmOttAudit 优酷OTT广告素材审核
// yunos.tvpubadmin.adm.ott.audit
//
// 审核优酷OTT端广告素材
func YunosTvpubadminAdmOttAudit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminAdmOttAuditAPIRequest, resp *tvupadmin.YunosTvpubadminAdmOttAuditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
