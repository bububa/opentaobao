package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaoorderpush 推送订单
// alibaba.ele.fengniao.order.push
//
// 推送淘宝订单至蜂鸟开放平台配送
func Alibabaelefengniaoorderpush(clt *core.SDKClient, req *logistic.AlibabaelefengniaoorderpushAPIRequest, session string) (*logistic.AlibabaelefengniaoorderpushAPIResponse, error) {
	var resp logistic.AlibabaelefengniaoorderpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
