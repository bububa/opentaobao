package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmactivityget 私域导购查询活动详情
// alibaba.lsy.crm.activity.get
//
// 私域导购查询活动详情
func Alibabalsycrmactivityget(clt *core.SDKClient, req *tmallnr.AlibabalsycrmactivitygetAPIRequest, session string) (*tmallnr.AlibabalsycrmactivitygetAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmactivitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
