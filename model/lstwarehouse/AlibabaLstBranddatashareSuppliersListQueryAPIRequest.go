package lstwarehouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstBranddatashareSuppliersListQueryAPIRequest
品牌数据授权的供应商列表 API请求
alibaba.lst.branddatashare.suppliers.list.query

品牌商查询品牌数据授权的供应商列表 */
type AlibabaLstBranddatashareSuppliersListQueryAPIRequest struct {
	model.Params
	// 入参
	_query *LstBmSupplierQuery
}

// New
