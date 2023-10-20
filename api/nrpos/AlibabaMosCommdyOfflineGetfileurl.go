package nrpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrpos"
)

// AlibabaMosCommdyOfflineGetfileurl 去前置机pos商品离线文件下载地址查询接口
// alibaba.mos.commdy.offline.getfileurl
//
// 去前置机-pos查询离线文件下载地址接口
func AlibabaMosCommdyOfflineGetfileurl(clt *core.SDKClient, req *nrpos.AlibabaMosCommdyOfflineGetfileurlAPIRequest, session string) (*nrpos.AlibabaMosCommdyOfflineGetfileurlAPIResponse, error) {
	var resp nrpos.AlibabaMosCommdyOfflineGetfileurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
