package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectadvisersubmit 提交楼盘顾问
// alibaba.alihouse.newhome.project.adviser.submit
//
// 提交楼盘顾问
func Alibabaalihousenewhomeprojectadvisersubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectadvisersubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectadvisersubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectadvisersubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
