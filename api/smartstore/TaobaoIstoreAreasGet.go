package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

// TaobaoIstoreAreasGet 智慧门店区域编码查询
// taobao.istore.areas.get
//
// 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
// &lt;a href=&#34;http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html&#34;&gt;http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html&lt;/a&gt;
func TaobaoIstoreAreasGet(clt *core.SDKClient, req *smartstore.TaobaoIstoreAreasGetAPIRequest, resp *smartstore.TaobaoIstoreAreasGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
