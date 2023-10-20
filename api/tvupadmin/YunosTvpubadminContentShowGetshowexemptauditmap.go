package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentshowgetshowexemptauditmap 迎客松批量查询节目某个牌照的免审状态
// yunos.tvpubadmin.content.show.getshowexemptauditmap
//
// 迎客松批量查询节目某个牌照的免审状态
func Yunostvpubadmincontentshowgetshowexemptauditmap(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentshowgetshowexemptauditmapAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentshowgetshowexemptauditmapAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
