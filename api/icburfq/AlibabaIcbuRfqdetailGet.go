package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

/* AlibabaIcbuRfqdetailGet
获取RFQ详情
alibaba.icbu.rfqdetail.get

查看RFQ的详情信息 */
func AlibabaIcbuRfqdetailGet(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqdetailGetAPIRequest, session string) (*icburfq.AlibabaIcbuRfqdetailGetAPIResponse, error) {
	var resp icburfq.AlibabaIcbuRfqdetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
