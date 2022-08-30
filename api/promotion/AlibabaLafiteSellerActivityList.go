package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaLafiteSellerActivityList 商家自运营活动列表
// alibaba.lafite.seller.activity.list
//
// 商家查询自己配置的活动列表
func AlibabaLafiteSellerActivityList(clt *core.SDKClient, req *promotion.AlibabaLafiteSellerActivityListAPIRequest, session string) (*promotion.AlibabaLafiteSellerActivityListAPIResponse, error) {
	var resp promotion.AlibabaLafiteSellerActivityListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
