package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoopenuidgetbymixnick 通过mixnick转换openuid
// taobao.openuid.get.bymixnick
//
// 通过mixnick转换openuid
func Taobaoopenuidgetbymixnick(clt *core.SDKClient, req *util.TaobaoopenuidgetbymixnickAPIRequest, session string) (*util.TaobaoopenuidgetbymixnickAPIResponse, error) {
	var resp util.TaobaoopenuidgetbymixnickAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
