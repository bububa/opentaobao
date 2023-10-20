package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeprojectquery 查询楼盘相关信息
// alibaba.alihouse.newhome.project.query
//
// 根据outerid查询楼盘相关信息
func Alibabaalihousenewhomeprojectquery(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeprojectqueryAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeprojectqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeprojectqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
