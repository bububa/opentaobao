package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizEslInfoGetAPIRequest 价签设备信息查询接口 API请求
// taobao.uscesl.biz.esl.info.get
//
// 价签设备信息查询接口
type TaobaoUsceslBizEslInfoGetAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 门店storeId
	_storeId int64
	// 商家ID
	_bizBrandKey string
}

// NewTaobaoUsceslBizEslInfoGetRequest 初始化TaobaoUsceslBizEslInfoGetAPIRequest对象
func NewTaobaoUsceslBizEslInfoGetRequest() *TaobaoUsceslBizEslInfoGetAPIRequest {
	return &TaobaoUsceslBizEslInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizEslInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.esl.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizEslInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizEslInfoGetAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// Get EslBarCode Getter
func (r TaobaoUsceslBizEslInfoGetAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// Set is StoreId Setter
// 门店storeId
func (r *TaobaoUsceslBizEslInfoGetAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoUsceslBizEslInfoGetAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is BizBrandKey Setter
// 商家ID
func (r *TaobaoUsceslBizEslInfoGetAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// Get BizBrandKey Getter
func (r TaobaoUsceslBizEslInfoGetAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}
