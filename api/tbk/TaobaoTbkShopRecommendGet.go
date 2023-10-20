package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkshoprecommendget 淘宝客-公用-店铺关联推荐
// taobao.tbk.shop.recommend.get
//
// 入参卖家id，可推荐与此店铺相关联的相关店铺。
func Taobaotbkshoprecommendget(clt *core.SDKClient, req *tbk.TaobaotbkshoprecommendgetAPIRequest, session string) (*tbk.TaobaotbkshoprecommendgetAPIResponse, error) {
	var resp tbk.TaobaotbkshoprecommendgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
