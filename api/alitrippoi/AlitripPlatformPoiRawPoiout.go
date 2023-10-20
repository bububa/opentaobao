package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// Alitripplatformpoirawpoiout 飞猪poi输出
// alitrip.platform.poi.raw.poiout
//
// 输出指定城市poi指定信息
func Alitripplatformpoirawpoiout(clt *core.SDKClient, req *alitrippoi.AlitripplatformpoirawpoioutAPIRequest, session string) (*alitrippoi.AlitripplatformpoirawpoioutAPIResponse, error) {
	var resp alitrippoi.AlitripplatformpoirawpoioutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
