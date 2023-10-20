package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaTianmaoLanpeiLogisticsMailno 阿里巴巴.天猫家装.揽配.物流.获取运单号
// alibaba.tianmao.lanpei.logistics.mailno
//
// 阿里巴巴.天猫家装.揽配.物流.获取运单号
func AlibabaTianmaoLanpeiLogisticsMailno(clt *core.SDKClient, req *ascp.AlibabaTianmaoLanpeiLogisticsMailnoAPIRequest, session string) (*ascp.AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse, error) {
	var resp ascp.AlibabaTianmaoLanpeiLogisticsMailnoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
