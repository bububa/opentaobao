package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarpttargetingtageffectget 获取定向效果报表数据
// taobao.simba.rpt.targetingtageffect.get
//
// 获取定向效果报表数据
func Taobaosimbarpttargetingtageffectget(clt *core.SDKClient, req *simba.TaobaosimbarpttargetingtageffectgetAPIRequest, session string) (*simba.TaobaosimbarpttargetingtageffectgetAPIResponse, error) {
	var resp simba.TaobaosimbarpttargetingtageffectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
