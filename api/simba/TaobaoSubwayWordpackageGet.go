package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaywordpackageget 获取词包列表
// taobao.subway.wordpackage.get
//
// 获取流量智选、捡漏词包等词包列表
func Taobaosubwaywordpackageget(clt *core.SDKClient, req *simba.TaobaosubwaywordpackagegetAPIRequest, session string) (*simba.TaobaosubwaywordpackagegetAPIResponse, error) {
	var resp simba.TaobaosubwaywordpackagegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
