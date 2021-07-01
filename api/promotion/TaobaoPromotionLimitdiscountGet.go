package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* TaobaoPromotionLimitdiscountGet
限时打折查询
taobao.promotion.limitdiscount.get

分页查询某个卖家的限时打折信息。每页20条数据，按照结束时间降序排列。也可指定某一个限时打折id查询唯一的限时打折信息。 */
func TaobaoPromotionLimitdiscountGet(clt *core.SDKClient, req *promotion.TaobaoPromotionLimitdiscountGetAPIRequest, session string) (*promotion.TaobaoPromotionLimitdiscountGetAPIResponse, error) {
	var resp promotion.TaobaoPromotionLimitdiscountGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
