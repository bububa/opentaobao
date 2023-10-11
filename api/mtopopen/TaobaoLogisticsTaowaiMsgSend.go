package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsTaowaiMsgSend 淘外包裹物流信息走淘宝发包裹状态通知接口
// taobao.logistics.taowai.msg.send
//
// 淘外包裹物流信息走淘宝发包裹状态通知接口
func TaobaoLogisticsTaowaiMsgSend(clt *core.SDKClient, req *mtopopen.TaobaoLogisticsTaowaiMsgSendAPIRequest, session string) (*mtopopen.TaobaoLogisticsTaowaiMsgSendAPIResponse, error) {
	var resp mtopopen.TaobaoLogisticsTaowaiMsgSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
