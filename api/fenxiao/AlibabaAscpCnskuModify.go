package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuModify 供应链中台货品修改接口
// alibaba.ascp.cnsku.modify
//
// 供应链中台货品修改接口
func AlibabaAscpCnskuModify(clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuModifyAPIRequest, resp *fenxiao.AlibabaAscpCnskuModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
