package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentadvertqueryschedule 广告牌照管控查询
// yunos.tvpubadmin.content.advert.queryschedule
//
// 广告牌照管控查询
func Yunostvpubadmincontentadvertqueryschedule(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentadvertqueryscheduleAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentadvertqueryscheduleAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentadvertqueryscheduleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
