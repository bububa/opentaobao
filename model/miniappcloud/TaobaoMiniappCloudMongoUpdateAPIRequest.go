package miniappcloud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappcloudmongoupdateAPIRequest 更新MongoDB中的数据 API请求
// taobao.miniapp.cloud.mongo.update
//
// 更新MongoDB中的数据
type TaobaominiappcloudmongoupdateAPIRequest struct {
	model.Params
	// 更新条件
	_filter string
	// MongoDB表名
	_collection string
	// 待写入的数据
	_record string
	// 要操作的环境，默认是测试环境
	_env string
}

// NewTaobaominiappcloudmongoupdateRequest 初始化TaobaominiappcloudmongoupdateAPIRequest对象
func NewTaobaominiappcloudmongoupdateRequest() *TaobaominiappcloudmongoupdateAPIRequest {
	return &TaobaominiappcloudmongoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappcloudmongoupdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.mongo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappcloudmongoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappcloudmongoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFilter is Filter Setter
// 更新条件
func (r *TaobaominiappcloudmongoupdateAPIRequest) SetFilter(_filter string) error {
	r._filter = _filter
	r.Set("filter", _filter)
	return nil
}

// GetFilter Filter Getter
func (r TaobaominiappcloudmongoupdateAPIRequest) GetFilter() string {
	return r._filter
}

// SetCollection is Collection Setter
// MongoDB表名
func (r *TaobaominiappcloudmongoupdateAPIRequest) SetCollection(_collection string) error {
	r._collection = _collection
	r.Set("collection", _collection)
	return nil
}

// GetCollection Collection Getter
func (r TaobaominiappcloudmongoupdateAPIRequest) GetCollection() string {
	return r._collection
}

// SetRecord is Record Setter
// 待写入的数据
func (r *TaobaominiappcloudmongoupdateAPIRequest) SetRecord(_record string) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r TaobaominiappcloudmongoupdateAPIRequest) GetRecord() string {
	return r._record
}

// SetEnv is Env Setter
// 要操作的环境，默认是测试环境
func (r *TaobaominiappcloudmongoupdateAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// GetEnv Env Getter
func (r TaobaominiappcloudmongoupdateAPIRequest) GetEnv() string {
	return r._env
}
