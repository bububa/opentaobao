package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Taobaologisticswmsgoodsinfosync WMS回传货品长宽高图片等信息
// taobao.logistics.wms.goods.info.sync
//
// WMS回传货品长宽高图片等信息
func Taobaologisticswmsgoodsinfosync(clt *core.SDKClient, req *logistic.TaobaologisticswmsgoodsinfosyncAPIRequest, session string) (*logistic.TaobaologisticswmsgoodsinfosyncAPIResponse, error) {
	var resp logistic.TaobaologisticswmsgoodsinfosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
