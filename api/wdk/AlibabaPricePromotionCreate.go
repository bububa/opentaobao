package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabapricepromotioncreate 营销档期活动创建
// alibaba.price.promotion.create
//
// 大润发-盒马帮提供新增创建营销活动
func Alibabapricepromotioncreate(clt *core.SDKClient, req *wdk.AlibabapricepromotioncreateAPIRequest, session string) (*wdk.AlibabapricepromotioncreateAPIResponse, error) {
	var resp wdk.AlibabapricepromotioncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
