package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// Alibabaretailvendingpricewhitelistremove 价格管控白名单去除
// alibaba.retail.vending.price.whitelist.remove
//
// 商家价格管控白名单去除
func Alibabaretailvendingpricewhitelistremove(clt *core.SDKClient, req *retail.AlibabaretailvendingpricewhitelistremoveAPIRequest, session string) (*retail.AlibabaretailvendingpricewhitelistremoveAPIResponse, error) {
	var resp retail.AlibabaretailvendingpricewhitelistremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
