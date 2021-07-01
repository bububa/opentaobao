package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* TaobaoJstAstrolabeStoreinventoryAdjust
后端商品库存占用调整接口
taobao.jst.astrolabe.storeinventory.adjust

当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。 */
func TaobaoJstAstrolabeStoreinventoryAdjust(clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest, session string) (*omniorder.TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse, error) {
	var resp omniorder.TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
