package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilliilogisticsdetailurlget 电子面单物流详情授权url获取
// cainiao.waybill.ii.logisticsdetail.url.get
//
// 获取电子面单物流详情授权访问的H5 url
func Cainiaowaybilliilogisticsdetailurlget(clt *core.SDKClient, req *waybill.CainiaowaybilliilogisticsdetailurlgetAPIRequest, session string) (*waybill.CainiaowaybilliilogisticsdetailurlgetAPIResponse, error) {
	var resp waybill.CainiaowaybilliilogisticsdetailurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
