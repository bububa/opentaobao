package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmShopListGet 获取商圈商户信息列表
// alibaba.westcrm.shop.list.get
//
// 获取商圈商户信息列表
func AlibabaWestcrmShopListGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmShopListGetAPIRequest, session string) (*westcrm.AlibabaWestcrmShopListGetAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmShopListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
