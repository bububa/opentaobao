package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// Taobaotopsecretappkeybilldetail 服务商解密账单查询
// taobao.top.secret.appkey.bill.detail
//
// 服务商解密账单查询,分页返回所有店铺的账单，每个店铺每天仅包含两条数据，当天产生的号租费 和 当天产生的通话费，仅对90天内的账单提供SLA保障。查询账单详情请使用taobao.top.secret.bill.detail接口。
func Taobaotopsecretappkeybilldetail(clt *core.SDKClient, req *tbtrade.TaobaotopsecretappkeybilldetailAPIRequest, session string) (*tbtrade.TaobaotopsecretappkeybilldetailAPIResponse, error) {
	var resp tbtrade.TaobaotopsecretappkeybilldetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
