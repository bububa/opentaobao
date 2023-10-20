package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectdynamicdelete 楼盘快讯删除
// alibaba.alihouse.newhome.project.dynamic.delete
//
// 楼盘快讯删除
func Alibabaalihousenewhomeprojectdynamicdelete(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectdynamicdeleteAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectdynamicdeleteAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectdynamicdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
