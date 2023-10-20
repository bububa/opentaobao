package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomereviewchangestatus 楼盘测评草稿状态同步
// alibaba.alihouse.newhome.review.changestatus
//
// 楼盘测评草稿状态更新
func Alibabaalihousenewhomereviewchangestatus(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomereviewchangestatusAPIRequest, session string) (*alihouse.AlibabaalihousenewhomereviewchangestatusAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomereviewchangestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
