package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* TaobaoDeliveryTemplateDelete
删除运费模板
taobao.delivery.template.delete

根据用户指定的模板ID删除指定的模板 */
func TaobaoDeliveryTemplateDelete(clt *core.SDKClient, req *logistic.TaobaoDeliveryTemplateDeleteAPIRequest, session string) (*logistic.TaobaoDeliveryTemplateDeleteAPIResponse, error) {
	var resp logistic.TaobaoDeliveryTemplateDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
