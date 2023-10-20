package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmservicechannelshortlinkcreate ECRM创建淘短链服务
// taobao.crm.service.channel.shortlink.create
//
// 可生成店铺宝贝、店铺首页、活动链接、订单链接等4种可呼起手机淘宝APP至对应页面的淘短链。
func Taobaocrmservicechannelshortlinkcreate(clt *core.SDKClient, req *crm.TaobaocrmservicechannelshortlinkcreateAPIRequest, session string) (*crm.TaobaocrmservicechannelshortlinkcreateAPIResponse, error) {
	var resp crm.TaobaocrmservicechannelshortlinkcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
