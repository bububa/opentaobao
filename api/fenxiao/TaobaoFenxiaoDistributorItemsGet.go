package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDistributorItemsGet 查询商品下载记录
// taobao.fenxiao.distributor.items.get
//
// 供应商查询分销商商品下载记录。
func TaobaoFenxiaoDistributorItemsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorItemsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDistributorItemsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
