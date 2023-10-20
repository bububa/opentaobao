package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductGroupGet 分组信息获取
// alibaba.icbu.product.group.get
//
// 分组信息获取
func AlibabaIcbuProductGroupGet(clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupGetAPIRequest, resp *icbu.AlibabaIcbuProductGroupGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
