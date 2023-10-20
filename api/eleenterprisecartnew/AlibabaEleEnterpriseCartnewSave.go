package eleenterprisecartnew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecartnew"
)

// AlibabaEleEnterpriseCartnewSave 新版创建购物车
// alibaba.ele.enterprise.cartnew.save
//
// 新版创建购物车
func AlibabaEleEnterpriseCartnewSave(clt *core.SDKClient, req *eleenterprisecartnew.AlibabaEleEnterpriseCartnewSaveAPIRequest, resp *eleenterprisecartnew.AlibabaEleEnterpriseCartnewSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
