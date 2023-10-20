package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaominiappmesssagereply 轻店铺下行回复接口
// taobao.miniapp.messsage.reply
//
// 外部 isv 调用该进口来进行轻店铺消息的回复
func Taobaominiappmesssagereply(clt *core.SDKClient, req *user.TaobaominiappmesssagereplyAPIRequest, session string) (*user.TaobaominiappmesssagereplyAPIResponse, error) {
	var resp user.TaobaominiappmesssagereplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
