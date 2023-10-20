package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectsubmit 提交楼盘信息
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
func Alibabaalihousenewhomeprojectsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
