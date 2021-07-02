package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminAdmOttQuery 优酷OTT端广告素材查询
// yunos.tvpubadmin.adm.ott.query
//
// 查询广告素材
func YunosTvpubadminAdmOttQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminAdmOttQueryAPIRequest, session string) (*tvupadmin.YunosTvpubadminAdmOttQueryAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminAdmOttQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
