package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacreativeschangedget 分页获取修改过的广告创意ID和修改时间
// taobao.simba.creatives.changed.get
//
// 分页获取修改过的广告创意ID和修改时间
func Taobaosimbacreativeschangedget(clt *core.SDKClient, req *simba.TaobaosimbacreativeschangedgetAPIRequest, session string) (*simba.TaobaosimbacreativeschangedgetAPIResponse, error) {
	var resp simba.TaobaosimbacreativeschangedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
