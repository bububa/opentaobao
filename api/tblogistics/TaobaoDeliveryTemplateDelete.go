package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplateDelete 删除运费模板
// taobao.delivery.template.delete
//
// 根据用户指定的模板ID删除指定的模板
func TaobaoDeliveryTemplateDelete(clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplateDeleteAPIRequest, session string) (*tblogistics.TaobaoDeliveryTemplateDeleteAPIResponse, error) {
	var resp tblogistics.TaobaoDeliveryTemplateDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
