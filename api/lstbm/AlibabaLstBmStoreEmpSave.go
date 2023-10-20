package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

// AlibabaLstBmStoreEmpSave 保存品牌商自有门店和内部业代之间的关系
// alibaba.lst.bm.store.emp.save
//
// 保存品牌商自有门店和内部业代之间的关系
func AlibabaLstBmStoreEmpSave(clt *core.SDKClient, req *lstbm.AlibabaLstBmStoreEmpSaveAPIRequest, resp *lstbm.AlibabaLstBmStoreEmpSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
