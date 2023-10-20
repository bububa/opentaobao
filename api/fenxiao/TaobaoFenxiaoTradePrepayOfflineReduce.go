package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaotradeprepayofflinereduce 渠道分销供应商上传线下流水预存款（减少）
// taobao.fenxiao.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
func Taobaofenxiaotradeprepayofflinereduce(clt *core.SDKClient, req *fenxiao.TaobaofenxiaotradeprepayofflinereduceAPIRequest, session string) (*fenxiao.TaobaofenxiaotradeprepayofflinereduceAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaotradeprepayofflinereduceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
