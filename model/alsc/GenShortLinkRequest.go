package alsc

import (
	"sync"
)

// GenShortLinkRequest 结构体
type GenShortLinkRequest struct {
	// 域名类型，按实际填写，目前映射关系：{   &#34;alipay&#34;: &#34;alipays://platformapi/startapp&#34;,   &#34;duanqu&#34;: &#34;https://m.duanqu.com&#34;,   &#34;eleH5&#34;: &#34;https://h5.ele.me&#34;,   &#34;eleMiniApp&#34;: &#34;eleme://miniapp&#34;,   &#34;elemeSns&#34;: &#34;eleme://sns_share_v2&#34;,   &#34;koubei&#34;: &#34;koubei://platformapi/startapp&#34;,   &#34;mod&#34;: &#34;https://tb.ele.me&#34;,   &#34;modPpe&#34;: &#34;https://ppe-tb.ele.me&#34;,   &#34;ppe-r&#34;: &#34;https://ppe-r.ele.me&#34;,   &#34;pt&#34;: &#34;https://pt.ele.me&#34;,   &#34;ptPre&#34;: &#34;https://ppe-pt.ele.me&#34;,   &#34;r&#34;: &#34;https://r.ele.me&#34; }
	DomainType string `json:"domain_type,omitempty" xml:"domain_type,omitempty"`
	// 域名后面的路径
	Path string `json:"path,omitempty" xml:"path,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
}

var poolGenShortLinkRequest = sync.Pool{
	New: func() any {
		return new(GenShortLinkRequest)
	},
}

// GetGenShortLinkRequest() 从对象池中获取GenShortLinkRequest
func GetGenShortLinkRequest() *GenShortLinkRequest {
	return poolGenShortLinkRequest.Get().(*GenShortLinkRequest)
}

// ReleaseGenShortLinkRequest 释放GenShortLinkRequest
func ReleaseGenShortLinkRequest(v *GenShortLinkRequest) {
	v.DomainType = ""
	v.Path = ""
	v.BizScene = ""
	poolGenShortLinkRequest.Put(v)
}
