package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectcooperationsubmit 提交KA合作楼盘
// alibaba.alihouse.newhome.project.cooperation.submit
//
// 提交KA合作楼盘
func Alibabaalihousenewhomeprojectcooperationsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectcooperationsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectcooperationsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectcooperationsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
