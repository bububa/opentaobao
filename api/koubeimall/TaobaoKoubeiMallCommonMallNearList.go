package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonMallNearList 根据POI查询附近商圈列表信息
// taobao.koubei.mall.common.mall.near.list
//
// 通过用户/终端设备地理位置POI信息，查询附近商圈信息
func TaobaoKoubeiMallCommonMallNearList(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonMallNearListAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonMallNearListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
