package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationGoodsOnoffline 上线/下线 体检产品
// alibaba.alihealth.examination.goods.onoffline
//
// 第三方体检机构对接钉钉体检中的产品 上线／下线
func AlibabaAlihealthExaminationGoodsOnoffline(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationGoodsOnofflineAPIRequest, resp *examination.AlibabaAlihealthExaminationGoodsOnofflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
