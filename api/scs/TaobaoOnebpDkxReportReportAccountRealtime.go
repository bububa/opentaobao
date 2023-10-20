package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxreportreportaccountrealtime 获取账户实时报表
// taobao.onebp.dkx.report.report.account.realtime
//
// 获取账户实时报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{     &#34;log_date_list&#34;: [  &#34;2021-09-23&#34;     ]   }
func Taobaoonebpdkxreportreportaccountrealtime(clt *core.SDKClient, req *scs.TaobaoonebpdkxreportreportaccountrealtimeAPIRequest, session string) (*scs.TaobaoonebpdkxreportreportaccountrealtimeAPIResponse, error) {
	var resp scs.TaobaoonebpdkxreportreportaccountrealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
