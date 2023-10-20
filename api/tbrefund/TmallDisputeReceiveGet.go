package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Tmalldisputereceiveget 天猫逆向纠纷查询
// tmall.dispute.receive.get
//
// 展示商家所有退款信息
func Tmalldisputereceiveget(clt *core.SDKClient, req *tbrefund.TmalldisputereceivegetAPIRequest, session string) (*tbrefund.TmalldisputereceivegetAPIResponse, error) {
	var resp tbrefund.TmalldisputereceivegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
