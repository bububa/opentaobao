package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaordcaligeniuslogisticspackagesnotice 物流多包裹通知
// taobao.rdc.aligenius.logistics.packages.notice
//
// 订单发货之后，如果订单拆包、补发、赠品等场景，需要将多余包裹信息触达消费者, 大促会降级
func Taobaordcaligeniuslogisticspackagesnotice(clt *core.SDKClient, req *logistic.TaobaordcaligeniuslogisticspackagesnoticeAPIRequest, session string) (*logistic.TaobaordcaligeniuslogisticspackagesnoticeAPIResponse, error) {
	var resp logistic.TaobaordcaligeniuslogisticspackagesnoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
