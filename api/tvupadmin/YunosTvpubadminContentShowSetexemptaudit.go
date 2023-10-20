package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowSetexemptaudit 迎客松节目设置免审开关
// yunos.tvpubadmin.content.show.setexemptaudit
//
// 迎客松节目设置免审开关
func YunosTvpubadminContentShowSetexemptaudit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowSetexemptauditAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowSetexemptauditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
