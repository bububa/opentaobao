package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Taobaosellercatslistupdate 更新卖家自定义类目
// taobao.sellercats.list.update
//
// 此API更新卖家店铺内自定义类目 &lt;br/&gt;注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
func Taobaosellercatslistupdate(clt *core.SDKClient, req *shop.TaobaosellercatslistupdateAPIRequest, session string) (*shop.TaobaosellercatslistupdateAPIResponse, error) {
	var resp shop.TaobaosellercatslistupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
