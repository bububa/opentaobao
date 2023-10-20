package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentCompanyInfoGet 获取厂牌信息
// xiami.content.company.info.get
//
// 获取厂牌信息
func XiamiContentCompanyInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentCompanyInfoGetAPIRequest, session string) (*xiamicontent.XiamiContentCompanyInfoGetAPIResponse, error) {
	var resp xiamicontent.XiamiContentCompanyInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
