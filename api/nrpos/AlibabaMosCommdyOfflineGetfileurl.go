package nrpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrpos"
)

// Alibabamoscommdyofflinegetfileurl 去前置机pos商品离线文件下载地址查询接口
// alibaba.mos.commdy.offline.getfileurl
//
// 去前置机-pos查询离线文件下载地址接口
func Alibabamoscommdyofflinegetfileurl(clt *core.SDKClient, req *nrpos.AlibabamoscommdyofflinegetfileurlAPIRequest, session string) (*nrpos.AlibabamoscommdyofflinegetfileurlAPIResponse, error) {
	var resp nrpos.AlibabamoscommdyofflinegetfileurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
