package kbalgo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kbalgo"
)

// Alibabakbalgoalscpoisget 百度批量获取本地poi接口
// alibaba.kbalgo.alscpois.get
//
// 接口用于百度方获取本地生活poi数据，分页获取。
func Alibabakbalgoalscpoisget(clt *core.SDKClient, req *kbalgo.AlibabakbalgoalscpoisgetAPIRequest, session string) (*kbalgo.AlibabakbalgoalscpoisgetAPIResponse, error) {
	var resp kbalgo.AlibabakbalgoalscpoisgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
