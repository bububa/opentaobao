package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenexpressinfoquery 配送公司信息查询接口
// taobao.qimen.expressinfo.query
//
// 配送公司信息查询
func Taobaoqimenexpressinfoquery(clt *core.SDKClient, req *qimen.TaobaoqimenexpressinfoqueryAPIRequest, session string) (*qimen.TaobaoqimenexpressinfoqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenexpressinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
