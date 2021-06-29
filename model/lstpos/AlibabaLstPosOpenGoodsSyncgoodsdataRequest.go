package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量同步接口(最多10条商品信息) API请求
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

// 初始化AlibabaLstPosOpenGoodsSyncgoodsdataRequest对象
func NewAlibabaLstPosOpenGoodsSyncgoodsdataRequest() *AlibabaLstPosOpenGoodsSyncgoodsdataRequest{
    return &AlibabaLstPosOpenGoodsSyncgoodsdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.goods.syncgoodsdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoodsDTOList Setter
// 商品对象列表
func (r *AlibabaLstPosOpenGoodsSyncgoodsdataRequest) SetGoodsDTOList(goodsDTOList []GoodsDto) error {
    r.goodsDTOList = goodsDTOList
    r.Set("goods_d_t_o_list", goodsDTOList)
    return nil
}

// GoodsDTOList Getter
func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetGoodsDTOList() []GoodsDto {
    return r.goodsDTOList
}
// UserId Setter
// 用户主账号Id
func (r *AlibabaLstPosOpenGoodsSyncgoodsdataRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaLstPosOpenGoodsSyncgoodsdataRequest) GetUserId() int64 {
    return r.userId
}
