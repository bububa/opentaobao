package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmexchangeactivitycreate 创建积分兑换活动
// taobao.crm.exchange.activity.create
//
// 创建针对积分兑换类型的活动
func Taobaocrmexchangeactivitycreate(clt *core.SDKClient, req *crm.TaobaocrmexchangeactivitycreateAPIRequest, session string) (*crm.TaobaocrmexchangeactivitycreateAPIResponse, error) {
	var resp crm.TaobaocrmexchangeactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
