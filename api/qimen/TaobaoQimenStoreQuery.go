package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstorequery 门店信息查询接口
// taobao.qimen.store.query
//
// 商家在ERP等系统中调用该接口，查询门店相关信息
func Taobaoqimenstorequery(clt *core.SDKClient, req *qimen.TaobaoqimenstorequeryAPIRequest, session string) (*qimen.TaobaoqimenstorequeryAPIResponse, error) {
	var resp qimen.TaobaoqimenstorequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
