package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminadmottaudit 优酷OTT广告素材审核
// yunos.tvpubadmin.adm.ott.audit
//
// 审核优酷OTT端广告素材
func Yunostvpubadminadmottaudit(clt *core.SDKClient, req *tvupadmin.YunostvpubadminadmottauditAPIRequest, session string) (*tvupadmin.YunostvpubadminadmottauditAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminadmottauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
