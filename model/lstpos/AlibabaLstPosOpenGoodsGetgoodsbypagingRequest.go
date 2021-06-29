package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询用户全量的门店域商品接口(每页最多20条) APIRequest
alibaba.lst.pos.open.goods.getgoodsbypaging

分页查询用户全量的门店域商品接口(每页最多20条)
*/
type AlibabaLstPosOpenGoodsGetgoodsbypagingRequest struct {
    model.Params

    // 当前页码
    page   int64 

    // 当前主账号userId
    userId   int64 

}

func NewAlibabaLstPosOpenGoodsGetgoodsbypagingRequest() *AlibabaLstPosOpenGoodsGetgoodsbypagingRequest{
    return &AlibabaLstPosOpenGoodsGetgoodsbypagingRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.goods.getgoodsbypaging"
}

func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaLstPosOpenGoodsGetgoodsbypagingRequest) GetUserId() int64 {
    return r.userId
}

