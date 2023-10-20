package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Alibabalatourstrategyissue 阿里巴巴权益发放接口
// alibaba.latour.strategy.issue
//
// 阿里巴巴权益平台权益发放接口
func Alibabalatourstrategyissue(clt *core.SDKClient, req *promotion.AlibabalatourstrategyissueAPIRequest, session string) (*promotion.AlibabalatourstrategyissueAPIResponse, error) {
	var resp promotion.AlibabalatourstrategyissueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
