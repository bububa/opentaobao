package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiOrderComplaintGet 投诉详情
// alibaba.happytrip.taxi.order.complaint.get
//
// 获取投诉处理进度详情
func AlibabaHappytripTaxiOrderComplaintGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderComplaintGetAPIRequest, resp *happytrip.AlibabaHappytripTaxiOrderComplaintGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
