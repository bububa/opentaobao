package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDataQuery 魔盒统计数据查询接口
// yunos.tvpubadmin.data.query
//
// 用于华数查询魔盒上的一些用户统计数据
func YunosTvpubadminDataQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDataQueryAPIRequest, resp *tvupadmin.YunosTvpubadminDataQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
