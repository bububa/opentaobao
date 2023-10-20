package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Taobaophoneitemexternalrecommend 话费选品能力外放
// taobao.phone.item.external.recommend
//
// 话费选品能力外放
func Taobaophoneitemexternalrecommend(clt *core.SDKClient, req *alicom.TaobaophoneitemexternalrecommendAPIRequest, session string) (*alicom.TaobaophoneitemexternalrecommendAPIResponse, error) {
	var resp alicom.TaobaophoneitemexternalrecommendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
