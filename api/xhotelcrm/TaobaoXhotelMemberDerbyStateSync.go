package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// Taobaoxhotelmemberderbystatesync 德比侧同步卡、券状态接口
// taobao.xhotel.member.derby.state.sync
//
// 德比侧同步卡、券状态接口
func Taobaoxhotelmemberderbystatesync(clt *core.SDKClient, req *xhotelcrm.TaobaoxhotelmemberderbystatesyncAPIRequest, session string) (*xhotelcrm.TaobaoxhotelmemberderbystatesyncAPIResponse, error) {
	var resp xhotelcrm.TaobaoxhotelmemberderbystatesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
