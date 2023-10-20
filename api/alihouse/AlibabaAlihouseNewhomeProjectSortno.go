package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectsortno 新房排序值同步
// alibaba.alihouse.newhome.project.sortno
//
// 新房排序值同步
func Alibabaalihousenewhomeprojectsortno(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectsortnoAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectsortnoAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectsortnoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
