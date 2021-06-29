package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量同步接口(最多10条商品信息) APIRequest
alibaba.lst.pos.open.goods.syncgoodsdata

门店商品批量同步接口(最多10条商品信息)
*/
type AlibabaLstPosOpenGoodsSyncgoodsdataRequest struct {
    model.Params

    // 商品对象列表
    goodsDTOList   []GoodsDto 

    // 用户主账号Id
    userId   int64 

}

func NewAlibabaLstPosOpenGoodsSyncgoodsdataRequest() *AlibabaLstPosOpenGoodsSyncgoodsdataRequest{
    return &AlibabaLstPosOpenGoodsSyncgoodsdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.goods.syncgoodsdata"
}

func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstPosOpenGoodsSyncgoodsdataRequest) SetGoodsDTOList(goodsDTOList []GoodsDto) error {
    r.goodsDTOList = goodsDTOList
    r.Set("goods_d_t_o_list", goodsDTOList)
    return nil
}

func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetGoodsDTOList() []GoodsDto {
    return r.goodsDTOList
}

func (r *AlibabaLstPosOpenGoodsSyncgoodsdataRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetUserId() int64 {
    return r.userId
}

