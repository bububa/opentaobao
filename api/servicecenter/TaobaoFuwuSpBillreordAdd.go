package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaofuwuspbillreordadd 内购服务确认单明细上传接口
// taobao.fuwu.sp.billreord.add
//
// isv能通过该接口上传确认单明细数据
func Taobaofuwuspbillreordadd(clt *core.SDKClient, req *servicecenter.TaobaofuwuspbillreordaddAPIRequest, session string) (*servicecenter.TaobaofuwuspbillreordaddAPIResponse, error) {
	var resp servicecenter.TaobaofuwuspbillreordaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
