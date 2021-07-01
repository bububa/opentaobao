package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

/* AlibabaSeakingImagetranslate
图片机器翻译
alibaba.seaking.imagetranslate

图片机器翻译 */
func AlibabaSeakingImagetranslate(clt *core.SDKClient, req *seaking.AlibabaSeakingImagetranslateAPIRequest, session string) (*seaking.AlibabaSeakingImagetranslateAPIResponse, error) {
	var resp seaking.AlibabaSeakingImagetranslateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
