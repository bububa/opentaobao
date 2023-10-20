package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectsalestime 楼盘销售时刻同步
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
func Alibabaalihousenewhomeprojectsalestime(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectsalestimeAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectsalestimeAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectsalestimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
