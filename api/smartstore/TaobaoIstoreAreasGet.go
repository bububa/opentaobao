package smartstore

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/smartstore"
)

/* 
智慧门店区域编码查询 
taobao.istore.areas.get

查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html">http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2016/index.html</a>
*/
func TaobaoIstoreAreasGet(clt *core.SDKClient, req *smartstore.TaobaoIstoreAreasGetAPIRequest, session string) (*smartstore.TaobaoIstoreAreasGetAPIResponse, error) {
    var resp smartstore.TaobaoIstoreAreasGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
