package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolactivityskuadd 商品池活动新增商品
// alibaba.retail.marketing.itempool.activity.sku.add
//
// 新增或更新商品池活动商品信息【同城零售】
func Alibabaretailmarketingitempoolactivityskuadd(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolactivityskuaddAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolactivityskuaddAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolactivityskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
