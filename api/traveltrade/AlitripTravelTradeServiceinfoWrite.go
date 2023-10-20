package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptraveltradeserviceinfowrite 订单服务信息写入接口
// alitrip.travel.trade.serviceinfo.write
//
// 订单服务信息写入接口
func Alitriptraveltradeserviceinfowrite(clt *core.SDKClient, req *traveltrade.AlitriptraveltradeserviceinfowriteAPIRequest, session string) (*traveltrade.AlitriptraveltradeserviceinfowriteAPIResponse, error) {
	var resp traveltrade.AlitriptraveltradeserviceinfowriteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
