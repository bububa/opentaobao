package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

/* AlibabaSeakingImagetranslateSubmit
提交图片翻译任务
alibaba.seaking.imagetranslate.submit

提交图片翻译任务 */
func AlibabaSeakingImagetranslateSubmit(clt *core.SDKClient, req *seaking.AlibabaSeakingImagetranslateSubmitAPIRequest, session string) (*seaking.AlibabaSeakingImagetranslateSubmitAPIResponse, error) {
	var resp seaking.AlibabaSeakingImagetranslateSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
