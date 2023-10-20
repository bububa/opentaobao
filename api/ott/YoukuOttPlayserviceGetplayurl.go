package ott

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ott"
)

// Youkuottplayservicegetplayurl 获取播放串地址
// youku.ott.playservice.getplayurl
//
// 获取播放串地址服务
func Youkuottplayservicegetplayurl(clt *core.SDKClient, req *ott.YoukuottplayservicegetplayurlAPIRequest, session string) (*ott.YoukuottplayservicegetplayurlAPIResponse, error) {
	var resp ott.YoukuottplayservicegetplayurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
