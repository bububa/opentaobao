package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingimagetranslatesubmit 提交图片翻译任务
// alibaba.seaking.imagetranslate.submit
//
// 提交图片翻译任务
func Alibabaseakingimagetranslatesubmit(clt *core.SDKClient, req *seaking.AlibabaseakingimagetranslatesubmitAPIRequest, session string) (*seaking.AlibabaseakingimagetranslatesubmitAPIResponse, error) {
	var resp seaking.AlibabaseakingimagetranslatesubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
