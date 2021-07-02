package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAdvertGettypes 获取广告位类型
// yunos.tvpubadmin.content.advert.gettypes
//
// 获取广告位类型
func YunosTvpubadminContentAdvertGettypes(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAdvertGettypesAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentAdvertGettypesAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentAdvertGettypesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
