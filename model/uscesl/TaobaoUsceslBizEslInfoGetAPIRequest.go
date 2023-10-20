package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizeslinfogetAPIRequest 价签设备信息查询接口 API请求
// taobao.uscesl.biz.esl.info.get
//
// 价签设备信息查询接口
type TaobaousceslbizeslinfogetAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 商家ID
	_bizBrandKey string
	// 门店storeId
	_storeId int64
}

// NewTaobaousceslbizeslinfogetRequest 初始化TaobaousceslbizeslinfogetAPIRequest对象
func NewTaobaousceslbizeslinfogetRequest() *TaobaousceslbizeslinfogetAPIRequest {
	return &TaobaousceslbizeslinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizeslinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.esl.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizeslinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizeslinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEslBarCode is EslBarCode Setter
// 价签条码
func (r *TaobaousceslbizeslinfogetAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// GetEslBarCode EslBarCode Getter
func (r TaobaousceslbizeslinfogetAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// SetBizBrandKey is BizBrandKey Setter
// 商家ID
func (r *TaobaousceslbizeslinfogetAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaousceslbizeslinfogetAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 门店storeId
func (r *TaobaousceslbizeslinfogetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaousceslbizeslinfogetAPIRequest) GetStoreId() int64 {
	return r._storeId
}
