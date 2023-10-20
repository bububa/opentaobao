package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentshowgetbyshowid 迎客松根据节目id获取节目元数据
// yunos.tvpubadmin.content.show.getbyshowid
//
// 迎客松根据节目id获取节目元数据
func Yunostvpubadmincontentshowgetbyshowid(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentshowgetbyshowidAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentshowgetbyshowidAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentshowgetbyshowidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
