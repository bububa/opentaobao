package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitemdiscountdeleteactivity 删除商品特价活动
// alibaba.wdk.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
func Alibabawdkmarketingitemdiscountdeleteactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitemdiscountdeleteactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitemdiscountdeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
