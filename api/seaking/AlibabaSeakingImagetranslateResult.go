package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingimagetranslateresult 获取图片翻译任务结果
// alibaba.seaking.imagetranslate.result
//
// 获取图片翻译任务结果
func Alibabaseakingimagetranslateresult(clt *core.SDKClient, req *seaking.AlibabaseakingimagetranslateresultAPIRequest, session string) (*seaking.AlibabaseakingimagetranslateresultAPIResponse, error) {
	var resp seaking.AlibabaseakingimagetranslateresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
