package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// Alibabashenjingcoreactivitygetappshowlist 获取神鲸活动列表
// alibaba.shenjing.core.activity.getappshowlist
//
// 获取神鲸活动列表
func Alibabashenjingcoreactivitygetappshowlist(clt *core.SDKClient, req *shenjing.AlibabashenjingcoreactivitygetappshowlistAPIRequest, session string) (*shenjing.AlibabashenjingcoreactivitygetappshowlistAPIResponse, error) {
	var resp shenjing.AlibabashenjingcoreactivitygetappshowlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
