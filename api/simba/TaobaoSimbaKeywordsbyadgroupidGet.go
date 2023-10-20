package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordsbyadgroupidget 取得一个推广组的所有关键词
// taobao.simba.keywordsbyadgroupid.get
//
// 取得一个推广组的所有关键词
func Taobaosimbakeywordsbyadgroupidget(clt *core.SDKClient, req *simba.TaobaosimbakeywordsbyadgroupidgetAPIRequest, session string) (*simba.TaobaosimbakeywordsbyadgroupidgetAPIResponse, error) {
	var resp simba.TaobaosimbakeywordsbyadgroupidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
