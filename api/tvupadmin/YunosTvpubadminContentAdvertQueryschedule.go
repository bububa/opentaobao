package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAdvertQueryschedule 广告牌照管控查询
// yunos.tvpubadmin.content.advert.queryschedule
//
// 广告牌照管控查询
func YunosTvpubadminContentAdvertQueryschedule(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAdvertQueryscheduleAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentAdvertQueryscheduleAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentAdvertQueryscheduleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
