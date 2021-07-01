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
type AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest struct {
    model.Params
    // 商品对象列表
    _goodsDTOList   []GoodsDto
    // 用户主账号Id
    _userId   int64
}

// 初始化AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest对象
func NewAlibabaLstPosOpenGoodsSyncgoodsdataRequest() *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest{
    return &AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.goods.syncgoodsdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoodsDTOList Setter
// 商品对象列表
func (r *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) SetGoodsDTOList(_goodsDTOList []GoodsDto) error {
    r._goodsDTOList = _goodsDTOList
    r.Set("goods_d_t_o_list", _goodsDTOList)
    return nil
}

// GoodsDTOList Getter
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetGoodsDTOList() []GoodsDto {
    return r._goodsDTOList
}
// UserId Setter
// 用户主账号Id
func (r *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetUserId() int64 {
    return r._userId
}
