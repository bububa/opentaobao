package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainpurchaseorderpayurl 火车票采购商接口-获取支付链接
// taobao.train.purchase.order.payurl
//
// 火车票采购商接口-获取支付链接
func Taobaotrainpurchaseorderpayurl(clt *core.SDKClient, req *train.TaobaotrainpurchaseorderpayurlAPIRequest, session string) (*train.TaobaotrainpurchaseorderpayurlAPIResponse, error) {
	var resp train.TaobaotrainpurchaseorderpayurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
