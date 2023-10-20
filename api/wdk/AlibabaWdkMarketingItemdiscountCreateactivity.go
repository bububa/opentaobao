package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitemdiscountcreateactivity 创建商品特价活动
// alibaba.wdk.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
func Alibabawdkmarketingitemdiscountcreateactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitemdiscountcreateactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingitemdiscountcreateactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitemdiscountcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
