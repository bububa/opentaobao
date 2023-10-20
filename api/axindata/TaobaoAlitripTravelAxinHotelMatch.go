package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// Taobaoalitriptravelaxinhotelmatch 酒店匹配接口-阿信
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
func Taobaoalitriptravelaxinhotelmatch(clt *core.SDKClient, req *axindata.TaobaoalitriptravelaxinhotelmatchAPIRequest, session string) (*axindata.TaobaoalitriptravelaxinhotelmatchAPIResponse, error) {
	var resp axindata.TaobaoalitriptravelaxinhotelmatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
