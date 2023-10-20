package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectecode 楼盘e码更新
// alibaba.alihouse.newhome.project.ecode
//
// 更新楼盘ecode
func Alibabaalihousenewhomeprojectecode(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectecodeAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectecodeAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectecodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
