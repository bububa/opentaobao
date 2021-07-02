package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoWeikeSubscinfoGet 需求订单查询接口
// taobao.weike.subscinfo.get
//
// 需求订单查询接口
func TaobaoWeikeSubscinfoGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeSubscinfoGetAPIRequest, session string) (*servicecenter.TaobaoWeikeSubscinfoGetAPIResponse, error) {
	var resp servicecenter.TaobaoWeikeSubscinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
