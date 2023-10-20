package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentshowgetbyshowlongid 迎客松根据节目longid获取节目元数据
// yunos.tvpubadmin.content.show.getbyshowlongid
//
// 迎客松根据节目longid获取节目元数据
func Yunostvpubadmincontentshowgetbyshowlongid(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentshowgetbyshowlongidAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentshowgetbyshowlongidAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentshowgetbyshowlongidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
