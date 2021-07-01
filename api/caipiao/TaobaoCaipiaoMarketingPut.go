package caipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

/* TaobaoCaipiaoMarketingPut
创建或修改商家送彩票活动
taobao.caipiao.marketing.put

卖家通过此接口新增或修改送彩票活动的配置，比如活动时间、活动的条件等。

店铺装修、宝贝详情页的宣导素材示例：
https://gw.alicdn.com/tfs/TB1_nOiSXXXXXbgXXXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1FZX6SXXXXXXzXFXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1z4t8SXXXXXckXpXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1BhqgSXXXXXcDXXXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1TYt9SXXXXXXAXFXXXXXXXXXX-750-280.png
https://gw.alicdn.com/tfs/TB1tzpNSXXXXXacXVXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1UXdxSXXXXXXsapXXXXXXXXXX-790-280.png
https://gw.alicdn.com/tfs/TB1_gV.SXXXXXbZXpXXXXXXXXXX-790-280.png */
func TaobaoCaipiaoMarketingPut(clt *core.SDKClient, req *caipiao.TaobaoCaipiaoMarketingPutAPIRequest, session string) (*caipiao.TaobaoCaipiaoMarketingPutAPIResponse, error) {
	var resp caipiao.TaobaoCaipiaoMarketingPutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
