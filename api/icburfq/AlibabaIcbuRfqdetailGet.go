package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// Alibabaicburfqdetailget 获取RFQ详情
// alibaba.icbu.rfqdetail.get
//
// 查看RFQ的详情信息
func Alibabaicburfqdetailget(clt *core.SDKClient, req *icburfq.AlibabaicburfqdetailgetAPIRequest, session string) (*icburfq.AlibabaicburfqdetailgetAPIResponse, error) {
	var resp icburfq.AlibabaicburfqdetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
