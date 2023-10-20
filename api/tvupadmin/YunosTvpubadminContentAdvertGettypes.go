package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentadvertgettypes 获取广告位类型
// yunos.tvpubadmin.content.advert.gettypes
//
// 获取广告位类型
func Yunostvpubadmincontentadvertgettypes(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentadvertgettypesAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentadvertgettypesAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentadvertgettypesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
