package healthnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/healthnr"
)

// AlibabaHealthNrLogisticsQuery 阿里健康新零售物流详情接口
// alibaba.health.nr.logistics.query
//
// 对阿里健康o2o对接的商户提供查询物流单详情的能力
func AlibabaHealthNrLogisticsQuery(clt *core.SDKClient, req *healthnr.AlibabaHealthNrLogisticsQueryAPIRequest, resp *healthnr.AlibabaHealthNrLogisticsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
