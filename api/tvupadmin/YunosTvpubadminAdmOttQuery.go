package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminAdmOttQuery 优酷OTT端广告素材查询
// yunos.tvpubadmin.adm.ott.query
//
// 查询广告素材
func YunosTvpubadminAdmOttQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminAdmOttQueryAPIRequest, resp *tvupadmin.YunosTvpubadminAdmOttQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
