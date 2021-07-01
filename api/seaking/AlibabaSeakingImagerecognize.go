package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

/* AlibabaSeakingImagerecognize
图片语种识别
alibaba.seaking.imagerecognize

图片语种识别 */
func AlibabaSeakingImagerecognize(clt *core.SDKClient, req *seaking.AlibabaSeakingImagerecognizeAPIRequest, session string) (*seaking.AlibabaSeakingImagerecognizeAPIResponse, error) {
	var resp seaking.AlibabaSeakingImagerecognizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
