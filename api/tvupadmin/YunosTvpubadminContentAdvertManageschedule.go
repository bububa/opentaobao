package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentadvertmanageschedule 广告牌照管控修改
// yunos.tvpubadmin.content.advert.manageschedule
//
// 广告牌照管控修改
func Yunostvpubadmincontentadvertmanageschedule(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentadvertmanagescheduleAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentadvertmanagescheduleAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentadvertmanagescheduleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
