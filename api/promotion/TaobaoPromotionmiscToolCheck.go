package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscToolCheck UMP工具检测
// taobao.promotionmisc.tool.check
//
// UMP工具检测。ISV通过该接口检测（通过taobao.ump.tool.add）创建的UMP工具（tool）是否符合规范，如果不符合，则返回错误信息和对应的解决方案的；工具检测通过后才可以提交工具审核邮件，提交工具审核时，需提供该接口的返回值。
func TaobaoPromotionmiscToolCheck(clt *core.SDKClient, req *promotion.TaobaoPromotionmiscToolCheckAPIRequest, session string) (*promotion.TaobaoPromotionmiscToolCheckAPIResponse, error) {
	var resp promotion.TaobaoPromotionmiscToolCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
