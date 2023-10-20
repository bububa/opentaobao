package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaoinventorymodify 阿里巴巴.天猫.aic库存.修改
// alibaba.tianmao.inventory.modify
//
// 阿里巴巴.天猫.aic库存.修改
func Alibabatianmaoinventorymodify(clt *core.SDKClient, req *ascp.AlibabatianmaoinventorymodifyAPIRequest, session string) (*ascp.AlibabatianmaoinventorymodifyAPIResponse, error) {
	var resp ascp.AlibabatianmaoinventorymodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
