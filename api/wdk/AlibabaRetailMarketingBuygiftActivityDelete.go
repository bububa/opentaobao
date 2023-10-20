package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivityDelete 删除单品买赠活动
// alibaba.retail.marketing.buygift.activity.delete
//
// 同城零售单品特价活动删除
func AlibabaRetailMarketingBuygiftActivityDelete(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse, error) {
	var resp wdk.AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
