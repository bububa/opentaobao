package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardTemplateNew 新增实体卡模板
// alibaba.fundplatform.card.template.new
//
// 该接口由制卡商实现，当新增一个实体卡模板的时候，需要调用该接口，通知制卡商同步新增卡模板信息。
func AlibabaFundplatformCardTemplateNew(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardTemplateNewAPIRequest, resp *fundplatform.AlibabaFundplatformCardTemplateNewAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
