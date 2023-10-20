package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqdetailGet 获取RFQ详情
// alibaba.icbu.rfqdetail.get
//
// 查看RFQ的详情信息
func AlibabaIcbuRfqdetailGet(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqdetailGetAPIRequest, resp *icburfq.AlibabaIcbuRfqdetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
