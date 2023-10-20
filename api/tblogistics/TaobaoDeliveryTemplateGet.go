package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplateGet 获取用户指定运费模板信息
// taobao.delivery.template.get
//
// 获取用户指定运费模板信息
func TaobaoDeliveryTemplateGet(clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplateGetAPIRequest, resp *tblogistics.TaobaoDeliveryTemplateGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
