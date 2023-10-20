package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptravelhotelticketproductproductupdate 产品批量变更通知
// alitrip.travel.hotelticket.product.productupdate
//
// 产品批量变更通知
func Alitriptravelhotelticketproductproductupdate(clt *core.SDKClient, req *traveltrade.AlitriptravelhotelticketproductproductupdateAPIRequest, session string) (*traveltrade.AlitriptravelhotelticketproductproductupdateAPIResponse, error) {
	var resp traveltrade.AlitriptravelhotelticketproductproductupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
