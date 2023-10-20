package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectdynamicsubmit 提交楼盘快讯
// alibaba.alihouse.newhome.project.dynamic.submit
//
// 提交楼盘快讯
func Alibabaalihousenewhomeprojectdynamicsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectdynamicsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectdynamicsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectdynamicsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
