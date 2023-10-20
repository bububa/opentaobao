package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaoinventoryquery 阿里巴巴.天猫.aic库存.查询
// alibaba.tianmao.inventory.query
//
// 阿里巴巴.天猫.aic库存.查询
func Alibabatianmaoinventoryquery(clt *core.SDKClient, req *ascp.AlibabatianmaoinventoryqueryAPIRequest, session string) (*ascp.AlibabatianmaoinventoryqueryAPIResponse, error) {
	var resp ascp.AlibabatianmaoinventoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
