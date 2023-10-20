package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// Xiamicontentcompanyinfoget 获取厂牌信息
// xiami.content.company.info.get
//
// 获取厂牌信息
func Xiamicontentcompanyinfoget(clt *core.SDKClient, req *xiamicontent.XiamicontentcompanyinfogetAPIRequest, session string) (*xiamicontent.XiamicontentcompanyinfogetAPIResponse, error) {
	var resp xiamicontent.XiamicontentcompanyinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
