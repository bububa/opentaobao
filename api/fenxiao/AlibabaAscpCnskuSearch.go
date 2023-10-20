package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuSearch 供应链中台货品搜索接口
// alibaba.ascp.cnsku.search
//
// 供应链中台货品搜索接口
func AlibabaAscpCnskuSearch(clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuSearchAPIRequest, resp *fenxiao.AlibabaAscpCnskuSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
