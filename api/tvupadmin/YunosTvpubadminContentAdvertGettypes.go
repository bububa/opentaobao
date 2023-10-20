package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAdvertGettypes 获取广告位类型
// yunos.tvpubadmin.content.advert.gettypes
//
// 获取广告位类型
func YunosTvpubadminContentAdvertGettypes(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAdvertGettypesAPIRequest, resp *tvupadmin.YunosTvpubadminContentAdvertGettypesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
