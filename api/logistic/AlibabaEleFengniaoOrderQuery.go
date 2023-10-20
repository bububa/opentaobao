package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaoorderquery 查询订单基本信息
// alibaba.ele.fengniao.order.query
//
// 查询订单基本信息
func Alibabaelefengniaoorderquery(clt *core.SDKClient, req *logistic.AlibabaelefengniaoorderqueryAPIRequest, session string) (*logistic.AlibabaelefengniaoorderqueryAPIResponse, error) {
	var resp logistic.AlibabaelefengniaoorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
