package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitemdiscountcreateactivity 创建商品特价活动
// alibaba.hm.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
func Alibabahmmarketingitemdiscountcreateactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitemdiscountcreateactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitemdiscountcreateactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitemdiscountcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
