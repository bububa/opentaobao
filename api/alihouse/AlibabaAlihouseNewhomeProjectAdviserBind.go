package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectadviserbind 置业顾问批量绑定楼盘
// alibaba.alihouse.newhome.project.adviser.bind
//
// 置业顾问批量绑定楼盘
func Alibabaalihousenewhomeprojectadviserbind(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectadviserbindAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectadviserbindAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectadviserbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
