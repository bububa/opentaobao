package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitemdiscountqueryactivity 查找特价活动
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
func Alibabahmmarketingitemdiscountqueryactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitemdiscountqueryactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitemdiscountqueryactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitemdiscountqueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
