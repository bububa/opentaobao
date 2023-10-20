package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectkanamequery 查询KA楼盘名称
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
func Alibabaalihousenewhomeprojectkanamequery(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectkanamequeryAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectkanamequeryAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectkanamequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
