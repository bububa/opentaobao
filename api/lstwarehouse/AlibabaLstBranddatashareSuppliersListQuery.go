package lstwarehouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

/* AlibabaLstBranddatashareSuppliersListQuery
品牌数据授权的供应商列表
alibaba.lst.branddatashare.suppliers.list.query

品牌商查询品牌数据授权的供应商列表 */
func AlibabaLstBranddatashareSuppliersListQuery(clt *core.SDKClient, req *lstwarehouse.AlibabaLstBranddatashareSuppliersListQueryAPIRequest, session string) (*lstwarehouse.AlibabaLstBranddatashareSuppliersListQueryAPIResponse, error) {
	var resp lstwarehouse.AlibabaLstBranddatashareSuppliersListQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
