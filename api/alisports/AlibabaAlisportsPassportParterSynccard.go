package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportpartersynccard 阿里体育-卡信息同步接口
// alibaba.alisports.passport.parter.synccard
//
// 运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中
func Alibabaalisportspassportpartersynccard(clt *core.SDKClient, req *alisports.AlibabaalisportspassportpartersynccardAPIRequest, session string) (*alisports.AlibabaalisportspassportpartersynccardAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportpartersynccardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
