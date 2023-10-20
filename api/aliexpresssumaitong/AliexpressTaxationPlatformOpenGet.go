package aliexpresssumaitong

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// AliexpressTaxationPlatformOpenGet 平台固定参数获取
// aliexpress.taxation.platform.open.get
//
// Aliexpress开放平台固定参数获取
func AliexpressTaxationPlatformOpenGet(clt *core.SDKClient, req *aliexpresssumaitong.AliexpressTaxationPlatformOpenGetAPIRequest, session string) (*aliexpresssumaitong.AliexpressTaxationPlatformOpenGetAPIResponse, error) {
	var resp aliexpresssumaitong.AliexpressTaxationPlatformOpenGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
