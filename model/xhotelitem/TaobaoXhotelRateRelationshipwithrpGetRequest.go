package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据gid查询卖家下所有的rpId APIRequest
taobao.xhotel.rate.relationshipwithrp.get

根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据
*/
type TaobaoXhotelRateRelationshipwithrpGetRequest struct {
    model.Params

    // 宝贝的gid
    gid   int64 

    // 页数，可根据此值展示某页的数据。不填默认为1
    pageNo   int64 

}

func NewTaobaoXhotelRateRelationshipwithrpGetRequest() *TaobaoXhotelRateRelationshipwithrpGetRequest{
    return &TaobaoXhotelRateRelationshipwithrpGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.relationshipwithrp.get"
}

func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateRelationshipwithrpGetRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetGid() int64 {
    return r.gid
}

func (r *TaobaoXhotelRateRelationshipwithrpGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetPageNo() int64 {
    return r.pageNo
}

