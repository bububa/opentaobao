package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* TaobaoLogisticsOnlineConfirm
确认发货通知接口
taobao.logistics.online.confirm

<br><font color='red'>仅在使用taobao.logistics.online.send 发货时未输入运单号的情况下，需要使用该接口补充填写运单号，来确认发货。<br>
确认发货的目的是让交易流程继续走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】。</font> */
func TaobaoLogisticsOnlineConfirm(clt *core.SDKClient, req *logistic.TaobaoLogisticsOnlineConfirmAPIRequest, session string) (*logistic.TaobaoLogisticsOnlineConfirmAPIResponse, error) {
	var resp logistic.TaobaoLogisticsOnlineConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
