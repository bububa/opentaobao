package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentshowsetexemptaudit 迎客松节目设置免审开关
// yunos.tvpubadmin.content.show.setexemptaudit
//
// 迎客松节目设置免审开关
func Yunostvpubadmincontentshowsetexemptaudit(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentshowsetexemptauditAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentshowsetexemptauditAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentshowsetexemptauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
