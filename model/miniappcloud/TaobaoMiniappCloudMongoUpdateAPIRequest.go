package miniappcloud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudMongoUpdateAPIRequest 更新MongoDB中的数据 API请求
// taobao.miniapp.cloud.mongo.update
//
// 更新MongoDB中的数据
type TaobaoMiniappCloudMongoUpdateAPIRequest struct {
	model.Params
	// MongoDB表名
	_collection string
	// 更新条件
	_filter string
	// 待写入的数据
	_record string
	// 要操作的环境，默认是测试环境
	_env string
}

// NewTaobaoMiniappCloudMongoUpdateRequest 初始化TaobaoMiniappCloudMongoUpdateAPIRequest对象
func NewTaobaoMiniappCloudMongoUpdateRequest() *TaobaoMiniappCloudMongoUpdateAPIRequest {
	return &TaobaoMiniappCloudMongoUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappCloudMongoUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.cloud.mongo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappCloudMongoUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Collection Setter
// MongoDB表名
func (r *TaobaoMiniappCloudMongoUpdateAPIRequest) SetCollection(_collection string) error {
	r._collection = _collection
	r.Set("collection", _collection)
	return nil
}

// Get Collection Getter
func (r TaobaoMiniappCloudMongoUpdateAPIRequest) GetCollection() string {
	return r._collection
}

// Set is Filter Setter
// 更新条件
func (r *TaobaoMiniappCloudMongoUpdateAPIRequest) SetFilter(_filter string) error {
	r._filter = _filter
	r.Set("filter", _filter)
	return nil
}

// Get Filter Getter
func (r TaobaoMiniappCloudMongoUpdateAPIRequest) GetFilter() string {
	return r._filter
}

// Set is Record Setter
// 待写入的数据
func (r *TaobaoMiniappCloudMongoUpdateAPIRequest) SetRecord(_record string) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// Get Record Getter
func (r TaobaoMiniappCloudMongoUpdateAPIRequest) GetRecord() string {
	return r._record
}

// Set is Env Setter
// 要操作的环境，默认是测试环境
func (r *TaobaoMiniappCloudMongoUpdateAPIRequest) SetEnv(_env string) error {
	r._env = _env
	r.Set("env", _env)
	return nil
}

// Get Env Getter
func (r TaobaoMiniappCloudMongoUpdateAPIRequest) GetEnv() string {
	return r._env
}
