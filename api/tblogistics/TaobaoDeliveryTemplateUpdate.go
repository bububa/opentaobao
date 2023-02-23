package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplateUpdate 修改运费模板
// taobao.delivery.template.update
//
// 修改运费模板
func TaobaoDeliveryTemplateUpdate(clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplateUpdateAPIRequest, session string) (*tblogistics.TaobaoDeliveryTemplateUpdateAPIResponse, error) {
	var resp tblogistics.TaobaoDeliveryTemplateUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
