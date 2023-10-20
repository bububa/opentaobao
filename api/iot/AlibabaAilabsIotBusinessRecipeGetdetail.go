package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Alibabaailabsiotbusinessrecipegetdetail 获取食谱详情
// alibaba.ailabs.iot.business.recipe.getdetail
//
// 获取食谱详情接口，获取ISV自己的食谱详情数据
func Alibabaailabsiotbusinessrecipegetdetail(clt *core.SDKClient, req *iot.AlibabaailabsiotbusinessrecipegetdetailAPIRequest, session string) (*iot.AlibabaailabsiotbusinessrecipegetdetailAPIResponse, error) {
	var resp iot.AlibabaailabsiotbusinessrecipegetdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
