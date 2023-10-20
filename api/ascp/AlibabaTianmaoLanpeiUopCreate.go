package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaolanpeiuopcreate 阿里巴巴.天猫家装.揽配.履约订单.创建
// alibaba.tianmao.lanpei.uop.create
//
// 阿里巴巴.天猫家装.揽配.履约订单.创建
func Alibabatianmaolanpeiuopcreate(clt *core.SDKClient, req *ascp.AlibabatianmaolanpeiuopcreateAPIRequest, session string) (*ascp.AlibabatianmaolanpeiuopcreateAPIResponse, error) {
	var resp ascp.AlibabatianmaolanpeiuopcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
