package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenDeliveryorderBatchcreateAnswer 发货单创建结果通知接口(批量)
// taobao.qimen.deliveryorder.batchcreate.answer
//
// WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
func TaobaoQimenDeliveryorderBatchcreateAnswer(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest, resp *qimen.TaobaoQimenDeliveryorderBatchcreateAnswerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
