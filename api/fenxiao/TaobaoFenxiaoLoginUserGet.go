package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaologinuserget 获取分销用户登录信息
// taobao.fenxiao.login.user.get
//
// 获取用户登录信息
func Taobaofenxiaologinuserget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaologinusergetAPIRequest, session string) (*fenxiao.TaobaofenxiaologinusergetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaologinusergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
