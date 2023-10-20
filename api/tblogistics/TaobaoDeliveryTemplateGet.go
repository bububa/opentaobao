package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoDeliveryTemplateGet 获取用户指定运费模板信息
// taobao.delivery.template.get
//
// 获取用户指定运费模板信息
func TaobaoDeliveryTemplateGet(clt *core.SDKClient, req *tblogistics.TaobaoDeliveryTemplateGetAPIRequest, session string) (*tblogistics.TaobaoDeliveryTemplateGetAPIResponse, error) {
	var resp tblogistics.TaobaoDeliveryTemplateGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
