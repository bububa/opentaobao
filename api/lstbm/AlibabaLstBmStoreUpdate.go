package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

// AlibabaLstBmStoreUpdate 修改品牌商自有门店数据
// alibaba.lst.bm.store.update
//
// 修改品牌商自有门店数据
func AlibabaLstBmStoreUpdate(clt *core.SDKClient, req *lstbm.AlibabaLstBmStoreUpdateAPIRequest, resp *lstbm.AlibabaLstBmStoreUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
