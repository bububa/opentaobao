package lstwarehouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

// Alibabalstbranddatasharesupplierslistquery 品牌数据授权的供应商列表
// alibaba.lst.branddatashare.suppliers.list.query
//
// 品牌商查询品牌数据授权的供应商列表
func Alibabalstbranddatasharesupplierslistquery(clt *core.SDKClient, req *lstwarehouse.AlibabalstbranddatasharesupplierslistqueryAPIRequest, session string) (*lstwarehouse.AlibabalstbranddatasharesupplierslistqueryAPIResponse, error) {
	var resp lstwarehouse.AlibabalstbranddatasharesupplierslistqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
