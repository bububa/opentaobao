package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsFengchaoMsgSend 丰巢走淘宝发包裹状态通知接口
// taobao.logistics.fengchao.msg.send
//
// 丰巢走淘宝发包裹状态通知接口
func TaobaoLogisticsFengchaoMsgSend(clt *core.SDKClient, req *mtopopen.TaobaoLogisticsFengchaoMsgSendAPIRequest, session string) (*mtopopen.TaobaoLogisticsFengchaoMsgSendAPIResponse, error) {
	var resp mtopopen.TaobaoLogisticsFengchaoMsgSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
