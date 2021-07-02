package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryItemadjust 库存占用调整接口
// taobao.jst.astrolabe.storeinventory.itemadjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
func TaobaoJstAstrolabeStoreinventoryItemadjust(clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest, session string) (*omniorder.TaobaoJstAstrolabeStoreinventoryItemadjustAPIResponse, error) {
	var resp omniorder.TaobaoJstAstrolabeStoreinventoryItemadjustAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
