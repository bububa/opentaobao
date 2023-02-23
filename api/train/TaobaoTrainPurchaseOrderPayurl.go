package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainPurchaseOrderPayurl 火车票采购商接口-获取支付链接
// taobao.train.purchase.order.payurl
//
// 火车票采购商接口-获取支付链接
func TaobaoTrainPurchaseOrderPayurl(clt *core.SDKClient, req *train.TaobaoTrainPurchaseOrderPayurlAPIRequest, session string) (*train.TaobaoTrainPurchaseOrderPayurlAPIResponse, error) {
	var resp train.TaobaoTrainPurchaseOrderPayurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
