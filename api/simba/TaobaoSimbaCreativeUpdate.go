package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacreativeupdate 修改创意
// taobao.simba.creative.update
//
// 更新一个创意的信息，可以设置创意标题、创意图片
func Taobaosimbacreativeupdate(clt *core.SDKClient, req *simba.TaobaosimbacreativeupdateAPIRequest, session string) (*simba.TaobaosimbacreativeupdateAPIResponse, error) {
	var resp simba.TaobaosimbacreativeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
