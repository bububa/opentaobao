package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// TaobaoLogisticsShunfengModifydataSave 顺丰小程序修改配送信息回传接口
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
func TaobaoLogisticsShunfengModifydataSave(clt *core.SDKClient, req *mtopopen.TaobaoLogisticsShunfengModifydataSaveAPIRequest, resp *mtopopen.TaobaoLogisticsShunfengModifydataSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
