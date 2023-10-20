package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomebrokerquery 根据外部经纪人ID查询
// alibaba.alihouse.existinghome.broker.query
//
// 根据外部经纪人ID查询
func Alibabaalihouseexistinghomebrokerquery(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomebrokerqueryAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomebrokerqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomebrokerqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
