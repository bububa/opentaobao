package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsOfficialaccountOnline 聚石塔公众号上线
// taobao.jst.sms.officialaccount.online
//
// 聚石塔公众号上线
func TaobaoJstSmsOfficialaccountOnline(clt *core.SDKClient, req *jst.TaobaoJstSmsOfficialaccountOnlineAPIRequest, session string) (*jst.TaobaoJstSmsOfficialaccountOnlineAPIResponse, error) {
	var resp jst.TaobaoJstSmsOfficialaccountOnlineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
