package aliexpresssumaitong

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// Aliexpresstaxationplatformopenget 平台固定参数获取
// aliexpress.taxation.platform.open.get
//
// Aliexpress开放平台固定参数获取
func Aliexpresstaxationplatformopenget(clt *core.SDKClient, req *aliexpresssumaitong.AliexpresstaxationplatformopengetAPIRequest, session string) (*aliexpresssumaitong.AliexpresstaxationplatformopengetAPIResponse, error) {
	var resp aliexpresssumaitong.AliexpresstaxationplatformopengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
