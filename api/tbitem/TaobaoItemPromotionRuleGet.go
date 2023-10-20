package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitempromotionruleget 获取商品已生效营销活动更新规则
// taobao.item.promotion.rule.get
//
// 获取商品已生效的更新规则信息，主要包含库存禁止修改，商品一口价禁止修改，库存减少锁定等规则生效信息
func Taobaoitempromotionruleget(clt *core.SDKClient, req *tbitem.TaobaoitempromotionrulegetAPIRequest, session string) (*tbitem.TaobaoitempromotionrulegetAPIResponse, error) {
	var resp tbitem.TaobaoitempromotionrulegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
