package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonmallnearlist 根据POI查询附近商圈列表信息
// taobao.koubei.mall.common.mall.near.list
//
// 通过用户/终端设备地理位置POI信息，查询附近商圈信息
func Taobaokoubeimallcommonmallnearlist(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonmallnearlistAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonmallnearlistAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonmallnearlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
