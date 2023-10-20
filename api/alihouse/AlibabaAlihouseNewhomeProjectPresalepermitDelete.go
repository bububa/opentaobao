package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectpresalepermitdelete 删除楼盘预售证
// alibaba.alihouse.newhome.project.presalepermit.delete
//
// 删除楼盘预售证信息
func Alibabaalihousenewhomeprojectpresalepermitdelete(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectpresalepermitdeleteAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectpresalepermitdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
