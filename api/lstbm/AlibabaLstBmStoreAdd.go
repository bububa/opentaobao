package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

// AlibabaLstBmStoreAdd 导入品牌商自有门店
// alibaba.lst.bm.store.add
//
// 导入品牌商自有门店
func AlibabaLstBmStoreAdd(clt *core.SDKClient, req *lstbm.AlibabaLstBmStoreAddAPIRequest, resp *lstbm.AlibabaLstBmStoreAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
