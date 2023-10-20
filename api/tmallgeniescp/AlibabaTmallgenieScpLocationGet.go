package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescplocationget 2-IBP查询CDC和RDC数据接口
// alibaba.tmallgenie.scp.location.get
//
// 天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口
func Alibabatmallgeniescplocationget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescplocationgetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescplocationgetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescplocationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
