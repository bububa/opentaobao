package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectpresalepermitsubmit 提交预售证
// alibaba.alihouse.newhome.project.presalepermit.submit
//
// 提交楼盘预售证信息
func Alibabaalihousenewhomeprojectpresalepermitsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectpresalepermitsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectpresalepermitsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectpresalepermitsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
