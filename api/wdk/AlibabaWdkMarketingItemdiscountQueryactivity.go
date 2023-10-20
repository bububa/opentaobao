package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitemdiscountqueryactivity 查找特价活动
// alibaba.wdk.marketing.itemdiscount.queryactivity
//
// 查找特价活动
func Alibabawdkmarketingitemdiscountqueryactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitemdiscountqueryactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitemdiscountqueryactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitemdiscountqueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
