package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplateUpdate 修改运费模板
// taobao.delivery.template.update
//
// 修改运费模板
func TaobaoDeliveryTemplateUpdate(clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplateUpdateAPIRequest, resp *tblogistics.TaobaoDeliveryTemplateUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
