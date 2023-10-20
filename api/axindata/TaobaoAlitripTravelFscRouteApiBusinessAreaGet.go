package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelfscrouteapibusinessareaget 获取业务区域
// taobao.alitrip.travel.fsc.route.api.business.area.get
//
// 获取业务区域
func Taobaoalitriptravelfscrouteapibusinessareaget(clt *core.SDKClient, req *axindata.TaobaoalitriptravelfscrouteapibusinessareagetAPIRequest, session string) (*axindata.TaobaoalitriptravelfscrouteapibusinessareagetAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelfscrouteapibusinessareagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
