package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoIstoreAreasGetAPIRequest
智慧门店区域编码查询 API请求
taobao.istore.areas.get

查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a> */
type TaobaoIstoreAreasGetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
	_fields string
}

// New
