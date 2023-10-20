package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasearchtagtemplateget 获取搜索人群TOP用户可添加人群信息
// taobao.simba.searchtagtemplate.get
//
// 获取搜索人群用户可添加人群信息
func Taobaosimbasearchtagtemplateget(clt *core.SDKClient, req *simba.TaobaosimbasearchtagtemplategetAPIRequest, session string) (*simba.TaobaosimbasearchtagtemplategetAPIResponse, error) {
	var resp simba.TaobaosimbasearchtagtemplategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
