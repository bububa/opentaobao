package wangwang

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wangwang"
)

// TaobaoAirislandKefuevalGet 客服评价详情接口_V2
// taobao.airisland.kefueval.get
//
// 获取买家对客服的服务评价
//
// 注意：
//
// 1. 请求超时[isp.top-remote-connection-timeout]或者数据过大错误[isp.runtime-max-limit]：因为某些帐号请求的数据会非常大，【需要通过减少请求时间范围避免该问题】
//
// 2. 时间范围：[now()-90d&lt;=btime ~ etime &lt;= now()-1d ] AND etime-btime &lt;=7d
//
// 3. 变更eval_recer：可空，返回脱敏的买家nick，如：摩天轮 -&gt; 摩**
//
// 4. 新增labelName：可空
func TaobaoAirislandKefuevalGet(ctx context.Context, clt *core.SDKClient, req *wangwang.TaobaoAirislandKefuevalGetAPIRequest, resp *wangwang.TaobaoAirislandKefuevalGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
