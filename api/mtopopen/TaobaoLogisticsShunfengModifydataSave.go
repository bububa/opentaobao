package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsShunfengModifydataSave 顺丰小程序修改配送信息回传接口
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
func TaobaoLogisticsShunfengModifydataSave(clt *core.SDKClient, req *mtopopen.TaobaoLogisticsShunfengModifydataSaveAPIRequest, session string) (*mtopopen.TaobaoLogisticsShunfengModifydataSaveAPIResponse, error) {
	var resp mtopopen.TaobaoLogisticsShunfengModifydataSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
