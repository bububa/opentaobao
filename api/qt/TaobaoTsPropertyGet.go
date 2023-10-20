package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// Taobaotspropertyget 淘宝服务属性查询
// taobao.ts.property.get
//
// 淘宝服务属性查询
func Taobaotspropertyget(clt *core.SDKClient, req *qt.TaobaotspropertygetAPIRequest, session string) (*qt.TaobaotspropertygetAPIResponse, error) {
	var resp qt.TaobaotspropertygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
