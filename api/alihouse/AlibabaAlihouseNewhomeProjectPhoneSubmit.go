package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectphonesubmit 提交楼盘电话
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
func Alibabaalihousenewhomeprojectphonesubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectphonesubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectphonesubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectphonesubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
