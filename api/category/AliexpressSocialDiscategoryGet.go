package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// Aliexpresssocialdiscategoryget 展示类目获取接口
// aliexpress.social.discategory.get
//
// AE展示类目获取接口
func Aliexpresssocialdiscategoryget(clt *core.SDKClient, req *category.AliexpresssocialdiscategorygetAPIRequest, session string) (*category.AliexpresssocialdiscategorygetAPIResponse, error) {
	var resp category.AliexpresssocialdiscategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
