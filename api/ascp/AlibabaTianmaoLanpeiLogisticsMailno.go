package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabatianmaolanpeilogisticsmailno 阿里巴巴.天猫家装.揽配.物流.获取运单号
// alibaba.tianmao.lanpei.logistics.mailno
//
// 阿里巴巴.天猫家装.揽配.物流.获取运单号
func Alibabatianmaolanpeilogisticsmailno(clt *core.SDKClient, req *ascp.AlibabatianmaolanpeilogisticsmailnoAPIRequest, session string) (*ascp.AlibabatianmaolanpeilogisticsmailnoAPIResponse, error) {
	var resp ascp.AlibabatianmaolanpeilogisticsmailnoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
