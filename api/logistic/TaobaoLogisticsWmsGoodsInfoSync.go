package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsWmsGoodsInfoSync WMS回传货品长宽高图片等信息
// taobao.logistics.wms.goods.info.sync
//
// WMS回传货品长宽高图片等信息
func TaobaoLogisticsWmsGoodsInfoSync(clt *core.SDKClient, req *logistic.TaobaoLogisticsWmsGoodsInfoSyncAPIRequest, session string) (*logistic.TaobaoLogisticsWmsGoodsInfoSyncAPIResponse, error) {
	var resp logistic.TaobaoLogisticsWmsGoodsInfoSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
