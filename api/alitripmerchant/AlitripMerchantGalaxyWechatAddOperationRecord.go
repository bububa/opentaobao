package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechataddoperationrecord 用户领取会员卡记录接口
// alitrip.merchant.galaxy.wechat.add.operation.record
//
// 用户领取会员卡记录接口
func Alitripmerchantgalaxywechataddoperationrecord(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechataddoperationrecordAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechataddoperationrecordAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechataddoperationrecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
