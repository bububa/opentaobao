package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptravelhotelticketproductproductupdatepush 产品批量变更推送通知
// alitrip.travel.hotelticket.product.productupdatepush
//
// 产品批量变更推送通知
func Alitriptravelhotelticketproductproductupdatepush(clt *core.SDKClient, req *traveltrade.AlitriptravelhotelticketproductproductupdatepushAPIRequest, session string) (*traveltrade.AlitriptravelhotelticketproductproductupdatepushAPIResponse, error) {
	var resp traveltrade.AlitriptravelhotelticketproductproductupdatepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
