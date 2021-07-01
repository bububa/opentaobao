package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstBranddatashareStockdataQueryAPIRequest
查询品牌商品实仓库存/周转效能 API请求
alibaba.lst.branddatashare.stockdata.query

品牌商查询授权供应商实仓库存数据 */
type AlibabaLstBranddatashareStockdataQueryAPIRequest struct {
	model.Params
	// 入参
	_param *BmSupplierStockDataParam
}

// New
