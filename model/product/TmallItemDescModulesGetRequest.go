package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品描述模块信息获取 APIRequest
tmall.item.desc.modules.get

商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
*/
type TmallItemDescModulesGetRequest struct {
    model.Params

    // 淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122
    catId   int64 

    // 商家主帐号id
    usrId   string 

}

func NewTmallItemDescModulesGetRequest() *TmallItemDescModulesGetRequest{
    return &TmallItemDescModulesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemDescModulesGetRequest) GetApiMethodName() string {
    return "tmall.item.desc.modules.get"
}

func (r TmallItemDescModulesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemDescModulesGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TmallItemDescModulesGetRequest) GetCatId() int64 {
    return r.catId
}

func (r *TmallItemDescModulesGetRequest) SetUsrId(usrId string) error {
    r.usrId = usrId
    r.Set("usr_id", usrId)
    return nil
}

func (r TmallItemDescModulesGetRequest) GetUsrId() string {
    return r.usrId
}

