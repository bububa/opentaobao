package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAdvertManageschedule 广告牌照管控修改
// yunos.tvpubadmin.content.advert.manageschedule
//
// 广告牌照管控修改
func YunosTvpubadminContentAdvertManageschedule(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAdvertManagescheduleAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentAdvertManagescheduleAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentAdvertManagescheduleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
