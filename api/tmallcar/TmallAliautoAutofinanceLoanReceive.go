package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoautofinanceloanreceive 接收放款结果通知
// tmall.aliauto.autofinance.loan.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户支用放款的通知结果
func Tmallaliautoautofinanceloanreceive(clt *core.SDKClient, req *tmallcar.TmallaliautoautofinanceloanreceiveAPIRequest, session string) (*tmallcar.TmallaliautoautofinanceloanreceiveAPIResponse, error) {
	var resp tmallcar.TmallaliautoautofinanceloanreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
