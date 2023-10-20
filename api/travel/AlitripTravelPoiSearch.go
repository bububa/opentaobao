package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Alitriptravelpoisearch POI信息查询
// alitrip.travel.poi.search
//
// POI信息查询，用于商品更新使用
func Alitriptravelpoisearch(clt *core.SDKClient, req *travel.AlitriptravelpoisearchAPIRequest, session string) (*travel.AlitriptravelpoisearchAPIResponse, error) {
	var resp travel.AlitriptravelpoisearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
