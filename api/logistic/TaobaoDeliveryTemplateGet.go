package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoDeliveryTemplateGet 获取用户指定运费模板信息
// taobao.delivery.template.get
//
// 获取用户指定运费模板信息
func TaobaoDeliveryTemplateGet(clt *core.SDKClient, req *logistic.TaobaoDeliveryTemplateGetAPIRequest, session string) (*logistic.TaobaoDeliveryTemplateGetAPIResponse, error) {
	var resp logistic.TaobaoDeliveryTemplateGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
