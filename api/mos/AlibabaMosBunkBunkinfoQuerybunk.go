package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosBunkBunkinfoQuerybunk 根据合同号查询铺位信息
// alibaba.mos.bunk.bunkinfo.querybunk
//
// 根据合同号查询铺位信息
func AlibabaMosBunkBunkinfoQuerybunk(clt *core.SDKClient, req *mos.AlibabaMosBunkBunkinfoQuerybunkAPIRequest, resp *mos.AlibabaMosBunkBunkinfoQuerybunkAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
