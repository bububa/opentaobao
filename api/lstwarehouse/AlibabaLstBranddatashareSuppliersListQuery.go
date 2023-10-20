package lstwarehouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

// AlibabaLstBranddatashareSuppliersListQuery 品牌数据授权的供应商列表
// alibaba.lst.branddatashare.suppliers.list.query
//
// 品牌商查询品牌数据授权的供应商列表
func AlibabaLstBranddatashareSuppliersListQuery(clt *core.SDKClient, req *lstwarehouse.AlibabaLstBranddatashareSuppliersListQueryAPIRequest, resp *lstwarehouse.AlibabaLstBranddatashareSuppliersListQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
