package eleenterprisecartnew

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

// AlibabaEleEnterpriseCartnewQuery 新版购物车查询
// alibaba.ele.enterprise.cartnew.query
//
// 新版购物车查询
func AlibabaEleEnterpriseCartnewQuery(ctx context.Context, clt *core.SDKClient, req *eleenterprisecartnew.AlibabaEleEnterpriseCartnewQueryAPIRequest, resp *eleenterprisecartnew.AlibabaEleEnterpriseCartnewQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
