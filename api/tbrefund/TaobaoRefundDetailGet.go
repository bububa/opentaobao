package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefunddetailget 退款详情页渲染
// taobao.refund.detail.get
//
// 退款详情页渲染
func Taobaorefunddetailget(clt *core.SDKClient, req *tbrefund.TaobaorefunddetailgetAPIRequest, session string) (*tbrefund.TaobaorefunddetailgetAPIResponse, error) {
	var resp tbrefund.TaobaorefunddetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
