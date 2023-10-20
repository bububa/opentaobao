package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaowirelesscontentcheck 无线开放内容检查
// taobao.wireless.content.check
//
// 无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 &lt;a href=&#34;https://help.aliyun.com/document_detail/70439.html&#34; target=&#34;blank&#34;&gt;阿里云内容安全&lt;/a&gt;
func Taobaowirelesscontentcheck(clt *core.SDKClient, req *util.TaobaowirelesscontentcheckAPIRequest, session string) (*util.TaobaowirelesscontentcheckAPIResponse, error) {
	var resp util.TaobaowirelesscontentcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
