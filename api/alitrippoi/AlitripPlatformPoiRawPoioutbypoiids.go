package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// AlitripPlatformPoiRawPoioutbypoiids 根据poiId输出飞猪poi数据
// alitrip.platform.poi.raw.poioutbypoiids
//
// 根据poiId输出飞猪poi数据
func AlitripPlatformPoiRawPoioutbypoiids(clt *core.SDKClient, req *alitrippoi.AlitripPlatformPoiRawPoioutbypoiidsAPIRequest, session string) (*alitrippoi.AlitripPlatformPoiRawPoioutbypoiidsAPIResponse, error) {
	var resp alitrippoi.AlitripPlatformPoiRawPoioutbypoiidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
