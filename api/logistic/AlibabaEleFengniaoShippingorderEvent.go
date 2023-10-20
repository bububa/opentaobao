package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaoshippingorderevent 查询运单事件信息
// alibaba.ele.fengniao.shippingorder.event
//
// 查询运单事件信息
func Alibabaelefengniaoshippingorderevent(clt *core.SDKClient, req *logistic.AlibabaelefengniaoshippingordereventAPIRequest, session string) (*logistic.AlibabaelefengniaoshippingordereventAPIResponse, error) {
	var resp logistic.AlibabaelefengniaoshippingordereventAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
