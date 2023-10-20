package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationgoodspublish 体检机构对接_商品发布／更新
// alibaba.alihealth.examination.goods.publish
//
// 体检机构对接_商品发布／更新
func Alibabaalihealthexaminationgoodspublish(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationgoodspublishAPIRequest, session string) (*examination.AlibabaalihealthexaminationgoodspublishAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationgoodspublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
