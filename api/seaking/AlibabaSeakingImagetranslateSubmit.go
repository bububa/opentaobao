package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingImagetranslateSubmit 提交图片翻译任务
// alibaba.seaking.imagetranslate.submit
//
// 提交图片翻译任务
func AlibabaSeakingImagetranslateSubmit(clt *core.SDKClient, req *seaking.AlibabaSeakingImagetranslateSubmitAPIRequest, resp *seaking.AlibabaSeakingImagetranslateSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
