package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoNextoneLogisticsSignUpdate AG物流签收状态写接口
// taobao.nextone.logistics.sign.update
//
// 商家上传退货的签收状态给AG
func TaobaoNextoneLogisticsSignUpdate(clt *core.SDKClient, req *logistic.TaobaoNextoneLogisticsSignUpdateAPIRequest, resp *logistic.TaobaoNextoneLogisticsSignUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
