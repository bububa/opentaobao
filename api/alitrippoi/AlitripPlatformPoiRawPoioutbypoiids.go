package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// Alitripplatformpoirawpoioutbypoiids 根据poiId输出飞猪poi数据
// alitrip.platform.poi.raw.poioutbypoiids
//
// 根据poiId输出飞猪poi数据
func Alitripplatformpoirawpoioutbypoiids(clt *core.SDKClient, req *alitrippoi.AlitripplatformpoirawpoioutbypoiidsAPIRequest, session string) (*alitrippoi.AlitripplatformpoirawpoioutbypoiidsAPIResponse, error) {
	var resp alitrippoi.AlitripplatformpoirawpoioutbypoiidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
