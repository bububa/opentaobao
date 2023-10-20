package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaocangpeiuopcreate 阿里巴巴.天猫家装.仓配.履约订单.创建
// alibaba.tianmao.cangpei.uop.create
//
// 阿里巴巴.天猫家装.仓配.履约订单.创建
func Alibabatianmaocangpeiuopcreate(clt *core.SDKClient, req *ascp.AlibabatianmaocangpeiuopcreateAPIRequest, session string) (*ascp.AlibabatianmaocangpeiuopcreateAPIResponse, error) {
	var resp ascp.AlibabatianmaocangpeiuopcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
