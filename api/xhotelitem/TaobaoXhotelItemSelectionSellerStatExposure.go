package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelitemselectionsellerstatexposure 选品曝光趋势
// taobao.xhotel.item.selection.seller.stat.exposure
//
// 用于提供给商家获取选品曝光趋势
func Taobaoxhotelitemselectionsellerstatexposure(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelitemselectionsellerstatexposureAPIRequest, session string) (*xhotelitem.TaobaoxhotelitemselectionsellerstatexposureAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelitemselectionsellerstatexposureAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
