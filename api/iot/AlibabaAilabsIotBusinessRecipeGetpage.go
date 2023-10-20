package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsiotbusinessrecipegetpage 分页查询食谱
// alibaba.ailabs.iot.business.recipe.getpage
//
// 分页查询食谱数据
func Alibabaailabsiotbusinessrecipegetpage(clt *core.SDKClient, req *iot.AlibabaailabsiotbusinessrecipegetpageAPIRequest, session string) (*iot.AlibabaailabsiotbusinessrecipegetpageAPIResponse, error) {
	var resp iot.AlibabaailabsiotbusinessrecipegetpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
