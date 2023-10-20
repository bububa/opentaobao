package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitemdiscountdeleteactivity 删除商品特价活动
// alibaba.hm.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
func Alibabahmmarketingitemdiscountdeleteactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitemdiscountdeleteactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitemdiscountdeleteactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitemdiscountdeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
