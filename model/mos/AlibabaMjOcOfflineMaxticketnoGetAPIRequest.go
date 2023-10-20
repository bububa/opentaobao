package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjocofflinemaxticketnogetAPIRequest pos机获取线下最大小票号 API请求
// alibaba.mj.oc.offline.maxticketno.get
//
// 给pos机提供线下最大小票号查询
type AlibabamjocofflinemaxticketnogetAPIRequest struct {
	model.Params
	// 日期
	_datetime string
	// 收银机号
	_posNo string
	// 外部门店号
	_storeNo string
}

// NewAlibabamjocofflinemaxticketnogetRequest 初始化AlibabamjocofflinemaxticketnogetAPIRequest对象
func NewAlibabamjocofflinemaxticketnogetRequest() *AlibabamjocofflinemaxticketnogetAPIRequest {
	return &AlibabamjocofflinemaxticketnogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjocofflinemaxticketnogetAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.offline.maxticketno.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjocofflinemaxticketnogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjocofflinemaxticketnogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDatetime is Datetime Setter
// 日期
func (r *AlibabamjocofflinemaxticketnogetAPIRequest) SetDatetime(_datetime string) error {
	r._datetime = _datetime
	r.Set("datetime", _datetime)
	return nil
}

// GetDatetime Datetime Getter
func (r AlibabamjocofflinemaxticketnogetAPIRequest) GetDatetime() string {
	return r._datetime
}

// SetPosNo is PosNo Setter
// 收银机号
func (r *AlibabamjocofflinemaxticketnogetAPIRequest) SetPosNo(_posNo string) error {
	r._posNo = _posNo
	r.Set("pos_no", _posNo)
	return nil
}

// GetPosNo PosNo Getter
func (r AlibabamjocofflinemaxticketnogetAPIRequest) GetPosNo() string {
	return r._posNo
}

// SetStoreNo is StoreNo Setter
// 外部门店号
func (r *AlibabamjocofflinemaxticketnogetAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabamjocofflinemaxticketnogetAPIRequest) GetStoreNo() string {
	return r._storeNo
}
