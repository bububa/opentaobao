package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// AlibabaAscpCnskuMappingDelete 货品关系解绑
// alibaba.ascp.cnsku.mapping.delete
//
// 货品关系解绑
func AlibabaAscpCnskuMappingDelete(clt *core.SDKClient, req *fenxiao.AlibabaAscpCnskuMappingDeleteAPIRequest, resp *fenxiao.AlibabaAscpCnskuMappingDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
