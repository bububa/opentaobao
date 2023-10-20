package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosubwaywordpackageupdate 批量更新词包
// taobao.subway.wordpackage.update
//
// 批量更新词包
func Taobaosubwaywordpackageupdate(clt *core.SDKClient, req *simba.TaobaosubwaywordpackageupdateAPIRequest, session string) (*simba.TaobaosubwaywordpackageupdateAPIResponse, error) {
	var resp simba.TaobaosubwaywordpackageupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
