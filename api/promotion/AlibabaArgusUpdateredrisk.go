package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaArgusUpdateredrisk 更新红线价格
// alibaba.argus.updateredrisk
//
// 商品健康中心新增红线价格规则
func AlibabaArgusUpdateredrisk(clt *core.SDKClient, req *promotion.AlibabaArgusUpdateredriskAPIRequest, resp *promotion.AlibabaArgusUpdateredriskAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
