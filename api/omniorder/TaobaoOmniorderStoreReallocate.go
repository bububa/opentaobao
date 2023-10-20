package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstorereallocate rellocate
// taobao.omniorder.store.reallocate
//
// 门店发货提供改派接口
func Taobaoomniorderstorereallocate(clt *core.SDKClient, req *omniorder.TaobaoomniorderstorereallocateAPIRequest, session string) (*omniorder.TaobaoomniorderstorereallocateAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstorereallocateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
