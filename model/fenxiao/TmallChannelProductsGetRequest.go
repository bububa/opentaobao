package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询供应商的产品数据 API请求
tmall.channel.products.get

查询供应商的产品数据。 

* 入参传入pids将优先查询，即只按这个条件查询。 
*入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
* 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
* 入参fields传入images将查询多图数据，不传只返回主图数据。 
* 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
* 查询结果按照产品发布时间倒序，即时间近的数据在前。
* 传入channel 渠道，会只返回相应渠道的产品
*/
type TmallChannelProductsGetRequest struct {
    model.Params
    // top_query_product_d_o
    _topQueryProductDO   *TopQueryProductDO
}

// 初始化TmallChannelProductsGetRequest对象
func NewTmallChannelProductsGetRequest() *TmallChannelProductsGetRequest{
    return &TmallChannelProductsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelProductsGetRequest) GetApiMethodName() string {
    return "tmall.channel.products.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopQueryProductDO Setter
// top_query_product_d_o
func (r *TmallChannelProductsGetRequest) SetTopQueryProductDO(_topQueryProductDO *TopQueryProductDO) error {
    r._topQueryProductDO = _topQueryProductDO
    r.Set("top_query_product_d_o", _topQueryProductDO)
    return nil
}

// TopQueryProductDO Getter
func (r TmallChannelProductsGetRequest) GetTopQueryProductDO() *TopQueryProductDO {
    return r._topQueryProductDO
}
