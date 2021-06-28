package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品列表 APIRequest
taobao.products.get

根据淘宝会员帐号搜索所有产品信息，推荐使用taobao.products.search
注意：支持分页，每页最多返回100条,默认值为40,页码从1开始，默认为第一页
*/
type TaobaoProductsGetRequest struct {
    model.Params

    // 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔
    fields   []string 

    // 用户昵称
    nick   string 

    // 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.
    pageNo   int64 

    // 每页条数.每页返回最多返回100条,默认值为40
    pageSize   int64 

}

func NewTaobaoProductsGetRequest() *TaobaoProductsGetRequest{
    return &TaobaoProductsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoProductsGetRequest) GetApiMethodName() string {
    return "taobao.products.get"
}

func (r TaobaoProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoProductsGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoProductsGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoProductsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoProductsGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoProductsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoProductsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoProductsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoProductsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

