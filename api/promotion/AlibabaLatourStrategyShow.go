package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabalatourstrategyshow 阿里巴巴权益投放接口
// alibaba.latour.strategy.show
//
// 阿里巴巴权益平台权益投放接口
func Alibabalatourstrategyshow(clt *core.SDKClient, req *promotion.AlibabalatourstrategyshowAPIRequest, session string) (*promotion.AlibabalatourstrategyshowAPIResponse, error) {
	var resp promotion.AlibabalatourstrategyshowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
