package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsOfficialaccountCancel 聚石塔取消公众号订购
// taobao.jst.sms.officialaccount.cancel
//
// 聚石塔取消公众号订购
func TaobaoJstSmsOfficialaccountCancel(clt *core.SDKClient, req *jst.TaobaoJstSmsOfficialaccountCancelAPIRequest, session string) (*jst.TaobaoJstSmsOfficialaccountCancelAPIResponse, error) {
	var resp jst.TaobaoJstSmsOfficialaccountCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
