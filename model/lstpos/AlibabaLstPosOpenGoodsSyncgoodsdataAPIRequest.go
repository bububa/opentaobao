package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopengoodssyncgoodsdataAPIRequest 门店商品批量同步接口(最多10条商品信息) API请求
// alibaba.lst.pos.open.goods.syncgoodsdata
//
// 门店商品批量同步接口(最多10条商品信息)
type AlibabalstposopengoodssyncgoodsdataAPIRequest struct {
	model.Params
	// 商品对象列表
	_goodsDTOList []GoodsDto
	// 用户主账号Id
	_userId int64
}

// NewAlibabalstposopengoodssyncgoodsdataRequest 初始化AlibabalstposopengoodssyncgoodsdataAPIRequest对象
func NewAlibabalstposopengoodssyncgoodsdataRequest() *AlibabalstposopengoodssyncgoodsdataAPIRequest {
	return &AlibabalstposopengoodssyncgoodsdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstposopengoodssyncgoodsdataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.goods.syncgoodsdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstposopengoodssyncgoodsdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstposopengoodssyncgoodsdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDTOList is GoodsDTOList Setter
// 商品对象列表
func (r *AlibabalstposopengoodssyncgoodsdataAPIRequest) SetGoodsDTOList(_goodsDTOList []GoodsDto) error {
	r._goodsDTOList = _goodsDTOList
	r.Set("goods_d_t_o_list", _goodsDTOList)
	return nil
}

// GetGoodsDTOList GoodsDTOList Getter
func (r AlibabalstposopengoodssyncgoodsdataAPIRequest) GetGoodsDTOList() []GoodsDto {
	return r._goodsDTOList
}

// SetUserId is UserId Setter
// 用户主账号Id
func (r *AlibabalstposopengoodssyncgoodsdataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalstposopengoodssyncgoodsdataAPIRequest) GetUserId() int64 {
	return r._userId
}
