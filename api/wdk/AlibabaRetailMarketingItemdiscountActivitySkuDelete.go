package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitemdiscountactivityskudelete 删除特价活动商品
// alibaba.retail.marketing.itemdiscount.activity.sku.delete
//
// 删除活动商品信息【同城零售】
func Alibabaretailmarketingitemdiscountactivityskudelete(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitemdiscountactivityskudeleteAPIRequest, session string) (*wdk.AlibabaretailmarketingitemdiscountactivityskudeleteAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitemdiscountactivityskudeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
