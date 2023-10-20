package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbacreativeidsdeletedget 获取删除的创意ID
// taobao.simba.creativeids.deleted.get
//
// 获取删除的创意ID
func Taobaosimbacreativeidsdeletedget(clt *core.SDKClient, req *simba.TaobaosimbacreativeidsdeletedgetAPIRequest, session string) (*simba.TaobaosimbacreativeidsdeletedgetAPIResponse, error) {
	var resp simba.TaobaosimbacreativeidsdeletedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
