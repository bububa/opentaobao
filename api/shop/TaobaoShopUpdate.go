package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Taobaoshopupdate 更新店铺基本信息
// taobao.shop.update
//
// 目前只支持标题、公告和描述的更新
func Taobaoshopupdate(clt *core.SDKClient, req *shop.TaobaoshopupdateAPIRequest, session string) (*shop.TaobaoshopupdateAPIResponse, error) {
	var resp shop.TaobaoshopupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
