package tmc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

/* TaobaoTmcUserGet
获取用户已开通消息
taobao.tmc.user.get

查询指定用户开通的消息通道和组 */
func TaobaoTmcUserGet(clt *core.SDKClient, req *tmc.TaobaoTmcUserGetAPIRequest, session string) (*tmc.TaobaoTmcUserGetAPIResponse, error) {
	var resp tmc.TaobaoTmcUserGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
