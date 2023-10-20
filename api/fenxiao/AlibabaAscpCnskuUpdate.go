package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuUpdate 供应链中台货品修改接口
// alibaba.ascp.cnsku.update
//
// 供应链中台货品修改接口
func AlibabaAscpCnskuUpdate(clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuUpdateAPIRequest, resp *fenxiao.AlibabaAscpCnskuUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
