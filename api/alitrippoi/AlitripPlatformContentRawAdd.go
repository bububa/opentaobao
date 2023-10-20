package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// AlitripPlatformContentRawAdd 穷游内容写入接口
// alitrip.platform.content.raw.add
//
// 穷游内容写入飞猪接口
func AlitripPlatformContentRawAdd(clt *core.SDKClient, req *alitrippoi.AlitripPlatformContentRawAddAPIRequest, resp *alitrippoi.AlitripPlatformContentRawAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
