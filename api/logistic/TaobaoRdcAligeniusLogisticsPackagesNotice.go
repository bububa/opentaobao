package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusLogisticsPackagesNotice 物流多包裹通知
// taobao.rdc.aligenius.logistics.packages.notice
//
// 订单发货之后，如果订单拆包、补发、赠品等场景，需要将多余包裹信息触达消费者, 大促会降级
func TaobaoRdcAligeniusLogisticsPackagesNotice(clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusLogisticsPackagesNoticeAPIRequest, session string) (*logistic.TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponse, error) {
	var resp logistic.TaobaoRdcAligeniusLogisticsPackagesNoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
