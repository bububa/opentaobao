package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】批量添加 APIRequest
taobao.alitrip.it.fare.batchadd

支持自有政策和销售规则批量添加，支持携程的数据格式。淘宝格式为list [object] to json string，object的属性和单条接口一致。每个接入方最多同时只能有1个处理中的导入任务，超过后直接返回失败。文件一定要zip压缩，压缩后大小不超过5M，编码格式utf-8
*/
type TaobaoAlitripItFareBatchaddRequest struct {
    model.Params

    // 新增类型，1 自有政策单程 2 自有政策往返 3 销售规则
    addType   int64 

    // 文本zip压缩后的数据字节流
    bytes   []*model.File 

    // 数据格式类型，1 淘宝 2 携程
    dataType   int64 

    // json格式的字符串，扩展属性，预留
    extendAttributes   string 

}

func NewTaobaoAlitripItFareBatchaddRequest() *TaobaoAlitripItFareBatchaddRequest{
    return &TaobaoAlitripItFareBatchaddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripItFareBatchaddRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.batchadd"
}

func (r TaobaoAlitripItFareBatchaddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripItFareBatchaddRequest) SetAddType(addType int64) error {
    r.addType = addType
    r.Set("addType", addType)
    return nil
}

func (r TaobaoAlitripItFareBatchaddRequest) GetAddType() int64 {
    return r.addType
}

func (r *TaobaoAlitripItFareBatchaddRequest) SetBytes(bytes []*model.File) error {
    r.bytes = bytes
    r.Set("bytes", bytes)
    return nil
}

func (r TaobaoAlitripItFareBatchaddRequest) GetBytes() []*model.File {
    return r.bytes
}

func (r *TaobaoAlitripItFareBatchaddRequest) SetDataType(dataType int64) error {
    r.dataType = dataType
    r.Set("dataType", dataType)
    return nil
}

func (r TaobaoAlitripItFareBatchaddRequest) GetDataType() int64 {
    return r.dataType
}

func (r *TaobaoAlitripItFareBatchaddRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

func (r TaobaoAlitripItFareBatchaddRequest) GetExtendAttributes() string {
    return r.extendAttributes
}

