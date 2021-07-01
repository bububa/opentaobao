package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

/* TaobaoTopAuthTokenRefresh
刷新Access Token
taobao.top.auth.token.refresh

根据refresh_token重新生成token，目前只有服务市场订购类应用可以刷新token，其他类型应用（如商家后台）使用固定时长token，不提供刷新功能。 */
func TaobaoTopAuthTokenRefresh(clt *core.SDKClient, req *util.TaobaoTopAuthTokenRefreshAPIRequest, session string) (*util.TaobaoTopAuthTokenRefreshAPIResponse, error) {
	var resp util.TaobaoTopAuthTokenRefreshAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
