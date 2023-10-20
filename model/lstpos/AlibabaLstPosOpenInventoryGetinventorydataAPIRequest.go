package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopeninventorygetinventorydataAPIRequest 商品库存只读接口(最多20条库存信息) API请求
// alibaba.lst.pos.open.inventory.getinventorydata
//
// 商品库存只读接口(最多20条库存信息)
type AlibabalstposopeninventorygetinventorydataAPIRequest struct {
	model.Params
	// ISV商品Id列表
	_isvGoodsIdList []string
	// 门店主账号Id
	_userId int64
}

// NewAlibabalstposopeninventorygetinventorydataRequest 初始化AlibabalstposopeninventorygetinventorydataAPIRequest对象
func NewAlibabalstposopeninventorygetinventorydataRequest() *AlibabalstposopeninventorygetinventorydataAPIRequest {
	return &AlibabalstposopeninventorygetinventorydataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstposopeninventorygetinventorydataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.inventory.getinventorydata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstposopeninventorygetinventorydataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstposopeninventorygetinventorydataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvGoodsIdList is IsvGoodsIdList Setter
// ISV商品Id列表
func (r *AlibabalstposopeninventorygetinventorydataAPIRequest) SetIsvGoodsIdList(_isvGoodsIdList []string) error {
	r._isvGoodsIdList = _isvGoodsIdList
	r.Set("isv_goods_id_list", _isvGoodsIdList)
	return nil
}

// GetIsvGoodsIdList IsvGoodsIdList Getter
func (r AlibabalstposopeninventorygetinventorydataAPIRequest) GetIsvGoodsIdList() []string {
	return r._isvGoodsIdList
}

// SetUserId is UserId Setter
// 门店主账号Id
func (r *AlibabalstposopeninventorygetinventorydataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalstposopeninventorygetinventorydataAPIRequest) GetUserId() int64 {
	return r._userId
}
