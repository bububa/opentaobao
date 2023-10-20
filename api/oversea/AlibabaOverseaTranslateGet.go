package oversea

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/oversea"
)

// Alibabaoverseatranslateget 获取文本翻译信息
// alibaba.oversea.translate.get
//
// 根据传入的文本信息，获取其目标语言的翻译结果
func Alibabaoverseatranslateget(clt *core.SDKClient, req *oversea.AlibabaoverseatranslategetAPIRequest, session string) (*oversea.AlibabaoverseatranslategetAPIResponse, error) {
	var resp oversea.AlibabaoverseatranslategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
