package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousecategorycontrolquery 类目权限查询
// alibaba.alihouse.category.control.query
//
// 类目权限查询
func Alibabaalihousecategorycontrolquery(clt *core.SDKClient, req *alihouse.AlibabaalihousecategorycontrolqueryAPIRequest, session string) (*alihouse.AlibabaalihousecategorycontrolqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihousecategorycontrolqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
