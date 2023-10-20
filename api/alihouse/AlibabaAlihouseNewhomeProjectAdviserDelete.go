package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectadviserdelete 删除楼盘顾问
// alibaba.alihouse.newhome.project.adviser.delete
//
// 删除楼盘顾问
func Alibabaalihousenewhomeprojectadviserdelete(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectadviserdeleteAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectadviserdeleteAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectadviserdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
