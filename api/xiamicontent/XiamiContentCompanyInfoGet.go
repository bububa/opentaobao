package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentCompanyInfoGet 获取厂牌信息
// xiami.content.company.info.get
//
// 获取厂牌信息
func XiamiContentCompanyInfoGet(clt *core.SDKClient, req *xiamicontent.XiamiContentCompanyInfoGetAPIRequest, resp *xiamicontent.XiamiContentCompanyInfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
