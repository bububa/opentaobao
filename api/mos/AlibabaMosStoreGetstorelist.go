package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosStoreGetstorelist 根据屏编号获取专柜集
// alibaba.mos.store.getstorelist
//
// 根据屏编号获取专柜集
func AlibabaMosStoreGetstorelist(clt *core.SDKClient, req *mos.AlibabaMosStoreGetstorelistAPIRequest, resp *mos.AlibabaMosStoreGetstorelistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
