package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// Alitripplatformpoirawfeed 存储poi原始数据
// alitrip.platform.poi.raw.feed
//
// 对接外部数据源，外部数据推送poi数据到飞猪
func Alitripplatformpoirawfeed(clt *core.SDKClient, req *alitrippoi.AlitripplatformpoirawfeedAPIRequest, session string) (*alitrippoi.AlitripplatformpoirawfeedAPIResponse, error) {
	var resp alitrippoi.AlitripplatformpoirawfeedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
