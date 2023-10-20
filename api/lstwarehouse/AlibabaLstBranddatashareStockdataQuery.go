package lstwarehouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

// AlibabaLstBranddatashareStockdataQuery 查询品牌商品实仓库存/周转效能
// alibaba.lst.branddatashare.stockdata.query
//
// 品牌商查询授权供应商实仓库存数据
func AlibabaLstBranddatashareStockdataQuery(clt *core.SDKClient, req *lstwarehouse.AlibabaLstBranddatashareStockdataQueryAPIRequest, resp *lstwarehouse.AlibabaLstBranddatashareStockdataQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
