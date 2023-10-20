package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplatesGet 获取用户下所有模板
// taobao.delivery.templates.get
//
// 根据用户ID获取用户下所有模板
func TaobaoDeliveryTemplatesGet(clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplatesGetAPIRequest, resp *tblogistics.TaobaoDeliveryTemplatesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
