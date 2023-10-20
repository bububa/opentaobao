package tuike

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuike"
)

// Alibabatuikeofferget 推广商品查询接口
// alibaba.tuike.offer.get
//
// 查询1688推客平台卖家推广中的商品信息
func Alibabatuikeofferget(clt *core.SDKClient, req *tuike.AlibabatuikeoffergetAPIRequest, session string) (*tuike.AlibabatuikeoffergetAPIResponse, error) {
	var resp tuike.AlibabatuikeoffergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
