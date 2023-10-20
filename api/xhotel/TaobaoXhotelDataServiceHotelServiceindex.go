package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// Taobaoxhoteldataservicehotelserviceindex 酒店服务指数
// taobao.xhotel.data.service.hotel.serviceindex
//
// 酒店服务指数
func Taobaoxhoteldataservicehotelserviceindex(clt *core.SDKClient, req *xhotel.TaobaoxhoteldataservicehotelserviceindexAPIRequest, session string) (*xhotel.TaobaoxhoteldataservicehotelserviceindexAPIResponse, error) {
	var resp xhotel.TaobaoxhoteldataservicehotelserviceindexAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
