package itpolicy

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/itpolicy"
)

/* 
【国际机票自有政策】批量添加 
taobao.alitrip.it.fare.batchadd

支持自有政策和销售规则批量添加，支持携程的数据格式。淘宝格式为list [object] to json string，object的属性和单条接口一致。每个接入方最多同时只能有1个处理中的导入任务，超过后直接返回失败。文件一定要zip压缩，压缩后大小不超过5M，编码格式utf-8
*/
func TaobaoAlitripItFareBatchadd(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareBatchaddRequest, session string) (*itpolicy.TaobaoAlitripItFareBatchaddAPIResponse, error) {
    var resp itpolicy.TaobaoAlitripItFareBatchaddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
