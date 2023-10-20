package eleenterprisecartnew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

// AlibabaEleEnterpriseCartnewQuery 新版购物车查询
// alibaba.ele.enterprise.cartnew.query
//
// 新版购物车查询
func AlibabaEleEnterpriseCartnewQuery(clt *core.SDKClient, req *eleenterprisecartnew.AlibabaEleEnterpriseCartnewQueryAPIRequest, resp *eleenterprisecartnew.AlibabaEleEnterpriseCartnewQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
