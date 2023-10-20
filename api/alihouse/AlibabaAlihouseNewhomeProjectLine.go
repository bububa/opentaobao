package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectline 楼盘上下架
// alibaba.alihouse.newhome.project.line
//
// 上下架楼盘
func Alibabaalihousenewhomeprojectline(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectlineAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectlineAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
