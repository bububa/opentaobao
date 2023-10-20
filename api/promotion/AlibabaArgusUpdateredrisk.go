package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabaargusupdateredrisk 更新红线价格
// alibaba.argus.updateredrisk
//
// 商品健康中心新增红线价格规则
func Alibabaargusupdateredrisk(clt *core.SDKClient, req *promotion.AlibabaargusupdateredriskAPIRequest, session string) (*promotion.AlibabaargusupdateredriskAPIResponse, error) {
	var resp promotion.AlibabaargusupdateredriskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
