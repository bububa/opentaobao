package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据gid查询卖家下所有的rpId API请求
taobao.xhotel.rate.relationshipwithrp.get

根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据
*/
type TaobaoXhotelRateRelationshipwithrpGetRequest struct {
    model.Params
    // 宝贝的gid
    _gid   int64
    // 页数，可根据此值展示某页的数据。不填默认为1
    _pageNo   int64
}

// 初始化TaobaoXhotelRateRelationshipwithrpGetRequest对象
func NewTaobaoXhotelRateRelationshipwithrpGetRequest() *TaobaoXhotelRateRelationshipwithrpGetRequest{
    return &TaobaoXhotelRateRelationshipwithrpGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.relationshipwithrp.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Gid Setter
// 宝贝的gid
func (r *TaobaoXhotelRateRelationshipwithrpGetRequest) SetGid(_gid int64) error {
    r._gid = _gid
    r.Set("gid", _gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetGid() int64 {
    return r._gid
}
// PageNo Setter
// 页数，可根据此值展示某页的数据。不填默认为1
func (r *TaobaoXhotelRateRelationshipwithrpGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoXhotelRateRelationshipwithrpGetRequest) GetPageNo() int64 {
    return r._pageNo
}
