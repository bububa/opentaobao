package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolactivityskudelete 删除商品池活动商品【同城零售】
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
func Alibabaretailmarketingitempoolactivityskudelete(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolactivityskudeleteAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolactivityskudeleteAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolactivityskudeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
