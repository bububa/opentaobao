package ott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ott"
)

// YoukuOttPlayserviceGetplayurl 获取播放串地址
// youku.ott.playservice.getplayurl
//
// 获取播放串地址服务
func YoukuOttPlayserviceGetplayurl(clt *core.SDKClient, req *ott.YoukuOttPlayserviceGetplayurlAPIRequest, resp *ott.YoukuOttPlayserviceGetplayurlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
