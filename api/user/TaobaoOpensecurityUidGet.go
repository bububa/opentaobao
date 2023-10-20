package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaoopensecurityuidget 淘宝open security uid 获取接口
// taobao.opensecurity.uid.get
//
// 根据明文 taobao user id 换取 app的 open_uid
func Taobaoopensecurityuidget(clt *core.SDKClient, req *user.TaobaoopensecurityuidgetAPIRequest, session string) (*user.TaobaoopensecurityuidgetAPIResponse, error) {
	var resp user.TaobaoopensecurityuidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
