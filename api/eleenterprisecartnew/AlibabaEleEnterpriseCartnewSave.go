package eleenterprisecartnew

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

// AlibabaEleEnterpriseCartnewSave 新版创建购物车
// alibaba.ele.enterprise.cartnew.save
//
// 新版创建购物车
func AlibabaEleEnterpriseCartnewSave(ctx context.Context, clt *core.SDKClient, req *eleenterprisecartnew.AlibabaEleEnterpriseCartnewSaveAPIRequest, resp *eleenterprisecartnew.AlibabaEleEnterpriseCartnewSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
