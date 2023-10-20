package lstpos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenInventoryGetinventorydataAPIRequest 商品库存只读接口(最多20条库存信息) API请求
// alibaba.lst.pos.open.inventory.getinventorydata
//
// 商品库存只读接口(最多20条库存信息)
type AlibabaLstPosOpenInventoryGetinventorydataAPIRequest struct {
	model.Params
	// ISV商品Id列表
	_isvGoodsIdList []string
	// 门店主账号Id
	_userId int64
}

// NewAlibabaLstPosOpenInventoryGetinventorydataRequest 初始化AlibabaLstPosOpenInventoryGetinventorydataAPIRequest对象
func NewAlibabaLstPosOpenInventoryGetinventorydataRequest() *AlibabaLstPosOpenInventoryGetinventorydataAPIRequest {
	return &AlibabaLstPosOpenInventoryGetinventorydataAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) Reset() {
	r._isvGoodsIdList = r._isvGoodsIdList[:0]
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.inventory.getinventorydata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvGoodsIdList is IsvGoodsIdList Setter
// ISV商品Id列表
func (r *AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) SetIsvGoodsIdList(_isvGoodsIdList []string) error {
	r._isvGoodsIdList = _isvGoodsIdList
	r.Set("isv_goods_id_list", _isvGoodsIdList)
	return nil
}

// GetIsvGoodsIdList IsvGoodsIdList Getter
func (r AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) GetIsvGoodsIdList() []string {
	return r._isvGoodsIdList
}

// SetUserId is UserId Setter
// 门店主账号Id
func (r *AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolAlibabaLstPosOpenInventoryGetinventorydataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstPosOpenInventoryGetinventorydataRequest()
	},
}

// GetAlibabaLstPosOpenInventoryGetinventorydataRequest 从 sync.Pool 获取 AlibabaLstPosOpenInventoryGetinventorydataAPIRequest
func GetAlibabaLstPosOpenInventoryGetinventorydataAPIRequest() *AlibabaLstPosOpenInventoryGetinventorydataAPIRequest {
	return poolAlibabaLstPosOpenInventoryGetinventorydataAPIRequest.Get().(*AlibabaLstPosOpenInventoryGetinventorydataAPIRequest)
}

// ReleaseAlibabaLstPosOpenInventoryGetinventorydataAPIRequest 将 AlibabaLstPosOpenInventoryGetinventorydataAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstPosOpenInventoryGetinventorydataAPIRequest(v *AlibabaLstPosOpenInventoryGetinventorydataAPIRequest) {
	v.Reset()
	poolAlibabaLstPosOpenInventoryGetinventorydataAPIRequest.Put(v)
}
