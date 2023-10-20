package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoautofinancecreditreceive 接收授信结果通知
// tmall.aliauto.autofinance.credit.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
func Tmallaliautoautofinancecreditreceive(clt *core.SDKClient, req *tmallcar.TmallaliautoautofinancecreditreceiveAPIRequest, session string) (*tmallcar.TmallaliautoautofinancecreditreceiveAPIResponse, error) {
	var resp tmallcar.TmallaliautoautofinancecreditreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
